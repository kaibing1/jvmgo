package heap

type Object struct {
	class *Class
	field Slots
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		field: newSlots(class.instanceSlotCount),
	}
}

func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.field
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}
