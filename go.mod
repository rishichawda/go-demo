module go_demo

go 1.15

replace demopackage => ./demo_package

replace examples => ./examples

require (
	demopackage v0.0.0-00010101000000-000000000000
	examples v0.0.0-00010101000000-000000000000 // indirect
)
