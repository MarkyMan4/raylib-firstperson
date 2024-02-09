package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	GRAVITY       = 0.1
	PLAYER_HEIGHT = 15
)

var (
	upVel           float32 = 0
	isOnGround              = true
	isCursorEnabled         = false
)

func updatePlayer(camera *rl.Camera3D) {
	// handle jumping
	if rl.IsKeyPressed(rl.KeySpace) && isOnGround {
		upVel += 2
		camera.Position.Y += upVel
		camera.Target.Y += upVel
		isOnGround = false
	}

	if camera.Position.Y > PLAYER_HEIGHT {
		camera.Position.Y += upVel
		camera.Target.Y += upVel
		upVel -= GRAVITY
	} else {
		camera.Position.Y = PLAYER_HEIGHT
		upVel = 0
		isOnGround = true
	}

	var xMoveFwd float32 = 0.0
	var xMoveBack float32 = 0.0
	var yMoveRight float32 = 0.0
	var yMoveLeft float32 = 0.0

	var moveSpeed float32 = 1

	if rl.IsKeyDown(rl.KeyLeftShift) {
		moveSpeed = 2
	}

	if rl.IsKeyDown(rl.KeyW) {
		xMoveFwd = moveSpeed
	}

	if rl.IsKeyDown(rl.KeyS) {
		xMoveBack = -moveSpeed
	}

	if rl.IsKeyDown(rl.KeyA) {
		yMoveLeft = -moveSpeed
	}

	if rl.IsKeyDown(rl.KeyD) {
		yMoveRight = moveSpeed
	}

	rl.UpdateCameraPro(
		camera,
		rl.Vector3{
			X: xMoveFwd + xMoveBack,
			Y: yMoveRight + yMoveLeft,
			Z: 0,
		},
		rl.Vector3{
			X: rl.GetMouseDelta().X * 0.1,
			Y: rl.GetMouseDelta().Y * 0.1,
			Z: 0,
		},
		0,
	)
}

func main() {
	var screenW int32 = 1500
	var screenH int32 = 800

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(screenW, screenH, "3d")

	// shader := rl.LoadShader("", "resources/textures/dream_vision.glsl")
	// target := rl.LoadRenderTexture(1500, 800)

	treeModel := rl.LoadModel("resources/models/Tree_01.obj")
	treeMaterial := rl.LoadTexture("resources/textures/green_stuff.jpg")
	treeModel.GetMaterials()[0].Maps.Texture = treeMaterial

	cubeModel := rl.LoadModelFromMesh(rl.GenMeshCube(10, 10, 10))
	woodTexture := rl.LoadTexture("resources/textures/default_wood.png")
	cubeModel.GetMaterials()[0].Maps.Texture = woodTexture

	grassCube := rl.LoadModelFromMesh(rl.GenMeshCube(10, 10, 10))
	grassTexture := rl.LoadTexture("resources/textures/grass.png")
	grassCube.GetMaterials()[0].Maps.Texture = grassTexture

	stoneCube := rl.LoadModelFromMesh(rl.GenMeshCube(10, 10, 10))
	stoneTexture := rl.LoadTexture("resources/textures/stone.png")
	stoneCube.GetMaterials()[0].Maps.Texture = stoneTexture

	enemySprite := rl.LoadTexture("assets/basic_enemy.png")

	camera := rl.Camera3D{
		Position:   rl.Vector3{X: 50.0, Y: 20.0, Z: 50.0},
		Target:     rl.Vector3{X: -50.0, Y: PLAYER_HEIGHT, Z: 0.0},
		Up:         rl.Vector3{X: 0.0, Y: 1.0, Z: 0.0},
		Fovy:       80.0,
		Projection: rl.CameraPerspective,
	}

	rl.SetTargetFPS(60)
	rl.DisableCursor()
	rl.SetExitKey(rl.KeyF4)

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyEscape) {
			if isCursorEnabled {
				rl.DisableCursor()
			} else {
				rl.EnableCursor()
			}

			isCursorEnabled = !isCursorEnabled
		}

		updatePlayer(&camera)

		rl.BeginDrawing()
		// rl.BeginTextureMode(target)
		rl.ClearBackground(rl.SkyBlue)
		rl.BeginMode3D(camera)
		rl.DrawPlane(rl.Vector3{X: camera.Position.X, Y: 0, Z: camera.Position.Z}, rl.Vector2{X: 500, Y: 500}, rl.DarkGreen)
		cubePos := rl.Vector3{X: 0.0, Y: 5, Z: 0.0}
		rl.DrawCube(cubePos, 10, 10, 10, rl.Blue)
		rl.DrawCubeWires(cubePos, 10, 10, 10, rl.Black)

		rl.DrawCapsule(rl.Vector3{X: 20, Y: 5, Z: 0}, rl.Vector3{X: 20, Y: 10, Z: 0}, 5, 20, 8, rl.Green)

		rl.DrawModel(treeModel, rl.Vector3{X: 0.0, Y: 7.5, Z: 50}, 20, rl.White)
		rl.DrawModel(cubeModel, rl.Vector3{X: 50.0, Y: 5, Z: 0.0}, 1, rl.White)
		rl.DrawModel(grassCube, rl.Vector3{X: 60.0, Y: 5, Z: 0.0}, 1, rl.White)
		rl.DrawModel(stoneCube, rl.Vector3{X: 70.0, Y: 5, Z: 0.0}, 1, rl.White)

		rl.DrawBillboard(camera, enemySprite, rl.Vector3{X: 0, Y: 10, Z: -30}, 20, rl.White)

		rl.EndMode3D()
		// rl.EndTextureMode()

		// rl.BeginDrawing()
		// rl.ClearBackground(rl.RayWhite)

		// rl.BeginShaderMode(shader)
		// rl.DrawTextureRec(target.Texture, rl.NewRectangle(0, 0, float32(target.Texture.Width), -float32(target.Texture.Height)), rl.NewVector2(0, 0), rl.White)
		// rl.EndShaderMode()

		// e.g. draw UI components
		// rl.DrawRectangle(
		// 	50, 50, 20, 20, rl.Red,
		// )

		// crosshair in center of screen
		rl.DrawCircle(int32(rl.GetScreenWidth())/2, int32(rl.GetScreenHeight())/2, 3, rl.Black)
		rl.DrawFPS(0, 0)

		rl.EndDrawing()
	}
}
