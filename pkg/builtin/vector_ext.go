package builtin

func ZeroVector2() Vector2 {
	return NewVector2WithFloat32Float32(0, 0)
}
func ZeroVector3() Vector3 {
	return NewVector3WithFloat32Float32Float32(0, 0, 0)
}
func ZeroVector4() Vector4 {
	return NewVector4WithFloat32Float32Float32Float32(0, 0, 0, 0)
}
func ZeroVector2i() Vector2i {
	return NewVector2iWithInt64Int64(0, 0)
}
func ZeroVector3i() Vector3i {
	return NewVector3iWithInt64Int64Int64(0, 0, 0)
}
func ZeroVector4i() Vector4i {
	return NewVector4iWithInt64Int64Int64Int64(0, 0, 0, 0)
}

// Vector2
func (cx Vector2) GetX() float32 {
	return cx.GetIndexed(0)
}
func (cx Vector2) SetX(value float32) {
	cx.SetIndexed(0, value)
}
func (cx Vector2) GetY() float32 {
	return cx.GetIndexed(1)
}
func (cx Vector2) SetY(value float32) {
	cx.SetIndexed(1, value)
}

// Vector3
func (cx Vector3) GetX() float32 {
	return cx.GetIndexed(0)
}
func (cx Vector3) SetX(value float32) {
	cx.SetIndexed(0, value)
}
func (cx Vector3) GetY() float32 {
	return cx.GetIndexed(1)
}
func (cx Vector3) SetY(value float32) {
	cx.SetIndexed(1, value)
}
func (cx Vector3) GetZ() float32 {
	return cx.GetIndexed(2)
}
func (cx Vector3) SetZ(value float32) {
	cx.SetIndexed(2, value)
}

// Vector4
func (cx Vector4) GetX() float32 {
	return cx.GetIndexed(0)
}
func (cx Vector4) SetX(value float32) {
	cx.SetIndexed(0, value)
}
func (cx Vector4) GetY() float32 {
	return cx.GetIndexed(1)
}
func (cx Vector4) SetY(value float32) {
	cx.SetIndexed(1, value)
}
func (cx Vector4) GetZ() float32 {
	return cx.GetIndexed(2)
}
func (cx Vector4) SetZ(value float32) {
	cx.SetIndexed(2, value)
}
func (cx Vector4) GetW() float32 {
	return cx.GetIndexed(3)
}
func (cx Vector4) SetW(value float32) {
	cx.SetIndexed(3, value)
}

// Vector2
func (cx Vector2i) GetX() int64 {
	return cx.GetIndexed(0)
}
func (cx Vector2i) SetX(value int64) {
	cx.SetIndexed(0, value)
}
func (cx Vector2i) GetY() int64 {
	return cx.GetIndexed(1)
}
func (cx Vector2i) SetY(value int64) {
	cx.SetIndexed(1, value)
}

// Vector3
func (cx Vector3i) GetX() int64 {
	return cx.GetIndexed(0)
}
func (cx Vector3i) SetX(value int64) {
	cx.SetIndexed(0, value)
}
func (cx Vector3i) GetY() int64 {
	return cx.GetIndexed(1)
}
func (cx Vector3i) SetY(value int64) {
	cx.SetIndexed(1, value)
}
func (cx Vector3i) GetZ() int64 {
	return cx.GetIndexed(2)
}
func (cx Vector3i) SetZ(value int64) {
	cx.SetIndexed(2, value)
}

// Vector4i
func (cx Vector4i) GetX() int64 {
	return cx.GetIndexed(0)
}
func (cx Vector4i) SetX(value int64) {
	cx.SetIndexed(0, value)
}
func (cx Vector4i) GetY() int64 {
	return cx.GetIndexed(1)
}
func (cx Vector4i) SetY(value int64) {
	cx.SetIndexed(1, value)
}
func (cx Vector4i) GetZ() int64 {
	return cx.GetIndexed(2)
}
func (cx Vector4i) SetZ(value int64) {
	cx.SetIndexed(2, value)
}
func (cx Vector4i) GetW() int64 {
	return cx.GetIndexed(3)
}
func (cx Vector4i) SetW(value int64) {
	cx.SetIndexed(3, value)
}
