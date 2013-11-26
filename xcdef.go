// Copyright (c) 2013 Cody Coons
//
// XCDEF Configuation Utility
package main

import (
  "os"
  "io"
  "github.com/codegangsta/cli"
)

func CreateGitIgnore() {
  gitIgnore := `# OS X Finder
.DS_Store

# Xcode per-user config
*.mode1
*.mode1v3
*.mode2v3
*.perspective
*.perspectivev3
*.pbxuser
xcuserdata
*.xccheckout

# Build products
build/
*.o
*.LinkFileList
*.hmap

# Automatic backup files
*~.nib/
*.swp
*~
*.dat
*.dep

# Cocoapods
Pods`

  f, err := os.Create(".gitignore")
  if err != nil {
    panic(err)
  }

  io.WriteString(f, gitIgnore)

  f.Close()
}

func CreateGitAttributes() {
  gitAttributes := "*.pbxproj binary merge=union"

  f, err := os.Create(".gitattributes")
  if err != nil {
    panic(err)
  }

  io.WriteString(f, gitAttributes)

  f.Close()
}

func main() {
  app := cli.NewApp()
  app.Name = "xcdef"
  app.Usage = "cli configuration utiltiy for XCode projects"

  app.Action = func(c *cli.Context) {
    CreateGitIgnore()
    println("** Created .gitignore")

    CreateGitAttributes()
    println("** Created .gitattributes")
  }

  app.Run(os.Args)
}
