package snowflake_ex

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)

// 魔改自 https://github.com/bwmarrin/snowflake/blob/master/snowflake.go
// 删除了一些没意义的代码
// 增加了等待至下一阶段时间周期的逻辑

var (
	_epoch     int64 = 1704038400000
	_nodeBits  uint8 = 10
	_stepBits  uint8 = 12
	_nodeMax   int64 = -1 ^ (-1 << _nodeBits)
	_nodeMod   int64 = 2 << (_nodeBits - 1)
	_nodeMask        = _nodeMax << _stepBits
	_stepMask  int64 = -1 ^ (-1 << _stepBits)
	_stepMax   int64 = 2 << (_stepBits - 1)
	_timeShift       = _nodeBits + _stepBits
	_nodeShift       = _stepBits
)

// A DistributedNode struct holds the basic information needed for a snowflake generator node
type DistributedNode struct {
	mu                        sync.Mutex
	epoch                     time.Time
	node                      int64
	time                      int64
	step                      int64
	forceBlockWhenAllConsumed bool // 当额度消费完之后，是否强制阻塞，等待下一轮时间周期
}

// An ID is a custom type used for a snowflake ID.  This is used so we can attach methods onto the ID.
type ID int64

// NewNode 0 <= node <= 1023
// forceBlockWhenAllConsumed 当额度消费完之后，是否强制阻塞，等待下一轮时间周期(高并发场景中建议设为true，避免id碰撞)
func NewNode(node uint32, forceBlockWhenAllConsumed bool) *DistributedNode {
	return &DistributedNode{
		mu:                        sync.Mutex{},
		epoch:                     time.Now().Add(time.Unix(_epoch/1000, (_epoch%1000)*1000000).Sub(time.Now())),
		node:                      int64(math.Abs(float64(node))) % _nodeMod,
		forceBlockWhenAllConsumed: forceBlockWhenAllConsumed,
	}
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (n *DistributedNode) Generate() ID {
	n.mu.Lock()
	defer n.mu.Unlock()

	cur := time.Since(n.epoch).Milliseconds()
	if n.forceBlockWhenAllConsumed {
		if (cur < n.time) || (cur == n.time && n.step >= _stepMax) {
			// 额度消费完之后，强制阻塞，等待下一轮时间周期
			time.Sleep(time.Duration(11-(cur%10)) * time.Millisecond)
		}
	}

	if cur > n.time { // 避免时间回拨
		// 进入下一轮时间周期
		n.time = cur
		n.step = 0
	}

	n.step += 1

	r := ID((cur)<<_timeShift |
		(n.node << _nodeShift) |
		(n.step),
	)

	return r
}

// Int64 returns an int64 of the snowflake ID
func (f ID) Int64() int64 {
	return int64(f)
}

// Uint64 returns an int64 of the snowflake ID
func (f ID) Uint64() uint64 {
	return uint64(f)
}

// String returns a string of the snowflake ID
func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10)
}

// ParseString converts a string into a snowflake ID
func ParseString(id string) (ID, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	return ID(i), err

}

// Base64 returns a base64 string of the snowflake ID
func (f ID) Base64() string {
	return base64.StdEncoding.EncodeToString(f.Bytes())
}

// Bytes returns a byte slice of the snowflake ID
func (f ID) Bytes() []byte {
	return []byte(f.String())
}

// IntBytes returns an array of bytes of the snowflake ID, encoded as a
// big endian integer.
func (f ID) IntBytes() [8]byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(f))
	return b
}

// Time returns an int64 unix timestamp in milliseconds of the snowflake ID time
// DEPRECATED: the below function will be removed in a future release.
func (f ID) Time() int64 {
	return (int64(f) >> _timeShift) + _epoch
}

// Node returns an int64 of the snowflake ID node number
func (f ID) Node() int64 {
	return int64(f) & _nodeMask >> _nodeShift
}

// Step returns an int64 of the snowflake step (or sequence) number
func (f ID) Step() int64 {
	return int64(f) & _stepMask
}

// MarshalJSON returns a json byte array string of the snowflake ID.
func (f ID) MarshalJSON() ([]byte, error) {
	buff := make([]byte, 0, 22)
	buff = append(buff, '"')
	buff = strconv.AppendInt(buff, int64(f), 10)
	buff = append(buff, '"')
	return buff, nil
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (f *ID) UnmarshalJSON(b []byte) error {
	if len(b) < 3 || b[0] != '"' || b[len(b)-1] != '"' {
		return errors.New(fmt.Sprintf("invalid snowflake ID %q", string(b)))
	}

	i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
	if err != nil {
		return err
	}

	*f = ID(i)
	return nil
}
