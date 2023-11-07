package typedeclimport

import subpkg "github.com/coderyw/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
