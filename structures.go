package SystemNX

/*   ______                           ____             _ __  _
   / ____/_  ________________  _____/ __ \____  _____(_) /_(_)___  ____
  / /   / / / / ___/ ___/ __ \/ ___/ /_/ / __ \/ ___/ / __/ / __ \/ __ \
 / /___/ /_/ / /  (__  ) /_/ / /  / ____/ /_/ (__  ) / /_/ / /_/ / / / /
 \____/\__,_/_/  /____/\____/_/  /_/    \____/____/_/\__/_/\____/_/ /_/

// --------- //
Fields:
-  X int32
-  Y int32


Methods:
-  GetX() int32
-  GetY() int32
*/

type CursorPosition struct {
	X, Y int32
}

func (position *CursorPosition) GetX() int32 {
	return position.X
}

func (position *CursorPosition) GetY() int32 {
	return position.Y
}

func (position *CursorPosition) SetX(x int32) {
	position.X = x
}

func (position *CursorPosition) SetY(y int32) {
	position.Y = y
}

func NewCursorPosition(X, Y int32) *CursorPosition {
	return &CursorPosition{X: X, Y: Y}
}
