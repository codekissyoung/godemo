package counters

// 未公开的类型值
type alertCounter int

// New 工厂函数
func New(value int) alertCounter {
	return alertCounter(value)
}
