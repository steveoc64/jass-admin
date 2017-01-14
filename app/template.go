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

      <label>Name</label>
      <input name="username" type="text" placeholder="User Name" autofocus>

      <label>Password</label>
      <input name="passwd" type="password" placeholder="Password">
      
      <!-- <input class="button button-outline" type="button" value="Cancel" id="l-cancelbtn"> -->
      <input id="sign-in" class="button button" type="submit" value="Sign In">
    </fieldset>
  </form>
</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("main-page", `<div class="jass-model-image"></div>

<div class="action-grid">
	<div class="action__item" url="/blog">
		<div class="action__title">Blogs</div>
		<div class="action__icon"><i class="fa fa-hashtag fa-lg"></i></div>
		<div class="action__text">
			Edit Blog posts.
		</div>
	</div>
	<div class="action__item" url="/category">
		<div class="action__title">Categories</div>
		<div class="action__icon"><i class="fa fa-cubes fa-lg"></i></div>
		<div class="action__text">
			Product Categories - Add / Edit / Delete.
		</div>
	</div>
	<div class="action__item" url="/product">
		<div class="action__title">Products</div>
		<div class="action__icon"><i class="fa fa-cube fa-lg"></i></div>
		<div class="action__text">
			Product Details - Add / Edit / Delete.
		</div>
	</div>
	<div class="action__item" url="/shipping">
		<div class="action__title">Shipping</div>
		<div class="action__icon"><i class="fa fa-ship fa-lg"></i></div>
		<div class="action__text">
			Shipping Constraints, Prices and Costs.
		</div>
	</div>
	<div class="action__item" url="/orders">
		<div class="action__title">Orders</div>
		<div class="action__icon"><i class="fa fa-vcard fa-lg"></i></div>
		<div class="action__text">
			Sales Orders, Tracking, Reconcilliation Reports.
		</div>
	</div>
	<div class="action__item" url="/newsletter">
		<div class="action__title">Newsletter</div>
		<div class="action__icon"><i class="fa fa-newspaper-o fa-lg"></i></div>
		<div class="action__text">
			Review and Generate Mailouts.
		</div>
	</div>
	<div class="action__item" url="/referrer">
		<div class="action__title">Referrers</div>
		<div class="action__icon"><i class="fa fa-cloud-download fa-lg"></i></div>
		<div class="action__text">
			Incoming Traffic from Referring Sites.
		</div>
	</div>
	<div class="action__item" url="/linkout">
		<div class="action__title">LinksOut</div>
		<div class="action__icon"><i class="fa fa-cloud-upload fa-lg"></i></div>
		<div class="action__text">
			Outgoing Traffic to External Sites.
		</div>
	</div>
	<div class="action__item" url="/customer">
		<div class="action__title">Customers</div>
		<div class="action__icon"><i class="fa fa-user fa-lg"></i></div>
		<div class="action__text">
			Customer Database.
		</div>
	</div>
	<div class="action__item" url="/users">
		<div class="action__title">Admins</div>
		<div class="action__icon"><i class="fa fa-user-secret fa-lg"></i></div>
		<div class="action__text">
			Admin Users.
		</div>
	</div>


</div>

<div class="jass-logo-small"></div>`); err != nil {
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
