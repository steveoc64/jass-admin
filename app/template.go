package main

// This package has been automatically generated with temple.
// Do not edit manually!

import (
	"github.com/go-humble/temple/temple"
)

var (
	GetTemplate     func(name string) (*temple.Template, error)
	GetPartial      func(name string) (*temple.Partial, error)
	GetLayout       func(name string) (*temple.Layout, error)
	MustGetTemplate func(name string) *temple.Template
	MustGetPartial  func(name string) *temple.Partial
	MustGetLayout   func(name string) *temple.Layout
)

func init() {
	var err error
	g := temple.NewGroup()

	if err = g.AddTemplate("login", `<!-- Login Screen -->
<div class="container" id="loginform">
  <form>
    <fieldset>

      <label for="l-username">Name</label>
      <input name="username" type="text" placeholder="User Name" id="l-username" autofocus>

      <label for="l-passwd">Password</label>
      <input name="passwd" type="password" placeholder="Password" id="l-passwd">
      
      <!-- <input class="button button-outline" type="button" value="Cancel" id="l-cancelbtn"> -->
      <input class="button button" type="submit" value="Sign In" id="l-loginbtn">
    </fieldset>
  </form>
</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("main-page", `<div class="container">
</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("slidemenu", `<!-- Slide in menu once logged in  -->
<nav class="cbp-spmenu cbp-spmenu-vertical cbp-spmenu-right" id="slidemenu">
  <a href="#" id="menu-fragrances"><i class="fa fa-snowflake-o"></i> Fragrances</a>
  <a href="#" id="menu-skincare"><i class="fa fa-hand-lizard-o"></i> Skincare</a>
  <a href="#" id="menu-merchandise"><i class="fa fa-gift"></i> Merchandise</a>
  <a href="#" id="menu-ambassadors"><i class="fa fa-user-circle-o"></i> Ambassadors</a>
  <a href="#" id="menu-blog"><i class="fa fa-hashtag"></i> Blog</a>
  <a href="#" id="menu-about"><i class="fa fa-question-circle-o"></i> About</a>
  <a href="#" id="menu-contact"><i class="fa fa-at"></i> Contact</a>
</nav> 



`); err != nil {
		panic(err)
	}

	GetTemplate = g.GetTemplate
	GetPartial = g.GetPartial
	GetLayout = g.GetLayout
	MustGetTemplate = g.MustGetTemplate
	MustGetPartial = g.MustGetPartial
	MustGetLayout = g.MustGetLayout
}
