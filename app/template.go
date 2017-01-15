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

	if err = g.AddTemplate("edit-form", `<div class="data-container">

  <form class="grid-form" action="#">
    <fieldset>
      <div class="row data-table-header">
        <h3 class="column column-90" id="legend">
          <i class="fa {{.Icon}} fa-lg" style="font-size: 3rem"></i>
          <span id="titletext">{{.Title}}</span>
        </h3>
        {{if .DeleteCB}}
        <div class="column col-center no-print">
          <i class="data-del-btn fa fa-minus-circle fa-lg"></i>    
        </div>
        {{end}}
        {{if .PrintCB}}
        <div class="column col-center no-print">
          <i class="data-print-btn fa fa-print fa-lg"></i>    
        </div>
        {{end}}        
      </div>

      {{range $rowindex,$row := .Rows}}
      <div data-row-span="{{.Span}}" name="row-{{$rowindex}}">
        {{range $fieldindex,$ := $row.Fields}}
        {{$fieldModel := .Model}}
        <div data-field-span="{{.Span}}" name="grid-{{$rowindex}}-{{$fieldindex}}">
          {{if ne .Type "button"}}
            {{if ne .Type "checkbox"}}
              <label>{{.Label}}</label>
            {{end}}
          {{end}}
          {{if eq .Type "div"}}
            <div name="{{.Model}}" class="{{.Class}}">Div Placeholder for {{.Model}}</div>
          {{end}}
          {{if eq .Type "swapper"}}
            {{$swapper := .Swapper}}
            {{range .Swapper.Panels}}
            <div class="swapper-option" name="{{$swapper.Name}}-{{.Name}}">
              {{$panelname := .Name}}
              {{range .Rows}}
              <div data-row-span="{{.Span}}">
                {{range .Fields}}
                  <div data-field-span="{{.Span}}">
                  {{if ne .Type "button"}}
                      <label>{{.Label}}</label>                  
                  {{end}}
                  {{if eq .Type "image"}}
                    <img src="/img/models/{{.Model}}">
                  {{end}}
                  {{if eq .Type "text"}}
                    <input type="{{.Type}}" name="{{$panelname}}-{{.Model}}" value="{{.Value}}" {{if .Readonly}}readonly{{end}} {{if .Autofocus}}autofocus{{end}}>
                  {{end}}                               
                  {{if eq .Type "date"}}
                    <input type="{{.Type}}" name="{{$panelname}}-{{.Model}}" value="{{.Value}}" {{if .Readonly}}readonly{{end}}>
                  {{end}}
                  {{if eq .Type "number"}}
                    <input type="{{.Type}}" name="{{$panelname}}-{{.Model}}" value="{{.Value}}" step="{{.Step}}" {{if .Readonly}}readonly{{end}}>
                  {{end}}
                  {{if eq .Type "textarea"}}
                    {{if .CodeBlock}}
                    <pre><code name="{{$panelname}}-{{.Model}}">{{.Value}}</code></pre>
                    {{else}}
                    <textarea name="{{$panelname}}-{{.Model}}" {{if .Readonly}}readonly{{end}} {{if .BigText}}class="bigtext"}}{{end}}>{{.Value}}</textarea>
                    {{end}}
                  {{end}}
                  {{if eq .Type "select"}}
                    {{if .Readonly}}
                      <input type="text" readonly name="{{$panelname}}-{{.Model}}" value="{{$.GetSelected}}">
                    {{else}}
                    <select name="{{$panelname}}-{{.Model}}">
                      {{range .Options}}
                        <option value="{{.Key}}" {{if .Selected}}selected{{end}}>{{.Display}}</option>
                      {{end}}
                    </select>
                    {{end}}
                  {{end}}
                  {{if eq .Type "checkbox"}}
                    <input type="checkbox" name="{{$panelname}}-{{.Model}}" {{if .Checked}}checked{{end}} {{if .Readonly}}disabled{{end}}>
                  {{end}}
                  {{if eq .Type "radio"}}
                    {{$model := .Model}}
                    <div name="{{$panelname}}-radio-{{$model}}">
                    <label></label>
                    {{range .Options}}
                    <label><input type="radio" name="{{$panelname}}-{{$model}}" value="{{.Key}}" {{if .Selected}}checked="checked"{{end}} {{if .Readonly}}disabled{{end}}> {{.Display}}</label>
                    {{end}}
                    </div>
                  {{end}}                  
                  {{if eq .Type "button"}}
                    <input type="button" class="button-outline" value="{{.Label}}" name="{{$panelname}}-{{.Model}}">
                    <!-- <button class="button-primary" name="{{$panelname}}-{{.Model}}">{{.Label}}</button>  -->
                  {{end}}
                  </div>
                {{end}}
              </div>
              {{end}}
            </div>
            {{end}}
          {{end}}
          {{if eq .Type "text"}}
            <input type="{{.Type}}" name="{{.Model}}" value="{{.Value}}" {{if .Focusme}}id="#focusme"{{end}} {{if .Readonly}}readonly{{end}} {{if .Autofocus}}autofocus{{end}}>
          {{end}}
          {{if eq .Type "photo"}}
            {{if .PhotoUpload}}
            <div class="image-upload">
              <span>
              <label for="file-input">
                <img src="/img/addPhoto.png">
              </label>
              <input id="file-input" type="file" accept="image/*;*.pdf" capture="camera" name="{{.Model}}" class="no-print"/><p>
              </span>
              <span>
                <img class="photouppreview hidden no-print" name="{{.Model}}Preview">
                <p class="previewhint hidden" name="{{.Model}}PreviewHint">Click on the preview to confirm upload ...</p>
              </span>
              <input name="{{.Model}}Filename" class="hidden" readonly>
            </div>
              <!-- <input type="file" name="{{.Model}}" multiple="multiple" class="no-print"><p> -->
              <!-- <img class="photouppreview hidden no-print" name="{{.Model}}-Preview"> -->
            {{end}}
            {{if .Preview}}
              <img class="photopreview hidden" name="{{.Model}}Preview"><br>
              <!-- <input name="{{.Model}}Filename" class="hidden" readonly><br> -->
              <span class="hidden" name="{{.Model}}Filename"></span>
            {{end}}
            {{if .Thumbnail}}
              <img class="photothumbnail hidden" name="{{.Model}}Preview">
            {{end}}
          {{end}}
          {{if eq .Type "image"}}
            <img src="{{.Value}}">
          {{end}}
          {{if eq .Type "date"}}
            <input type="{{.Type}}" name="{{.Model}}" value="{{.Value}}" {{if .Readonly}}readonly{{end}}>
          {{end}}
          {{if eq .Type "number"}}
            <input type="{{.Type}}" name="{{.Model}}" value="{{.Value}}" step="{{.Step}}" {{if .Readonly}}readonly{{end}}>
          {{end}}
          {{if eq .Type "textarea"}}
            {{if .CodeBlock}}
            <pre><code name="{{.Model}}">{{.Value}}</code></pre>
            {{else}}
            <textarea name="{{.Model}}" {{if .Readonly}}readonly{{end}} {{if .BigText}}class="bigtext"{{end}}>{{.Value}}</textarea>
            {{end}}
          {{end}}
          {{if eq .Type "select"}}
            {{if .Readonly}}
              <input type="text" readonly name="{{.Model}}" value="{{$.GetSelected}}">
            {{else}}
            <select name="{{.Model}}">
              {{range .Options}}
                <option value="{{.Key}}" {{if .Selected}}selected{{end}}>{{.Display}}</option>
              {{end}}
            </select>
            {{end}}
          {{end}}
          {{if eq .Type "groupselect"}}
            {{$sel := .Selected}}
            <select name="{{.Model}}">
              {{range .Group}}
              <optgroup {{if .Title}}label="{{.Title}}"{{end}}>
                {{range .Options}}
                <option value="{{.ID}}" {{if eq .ID $sel}}selected{{end}}>{{.Name}}</option>
                {{end}}
              </optgroup>
              {{end}}
            </select>
          {{end}}
          {{if eq .Type "checkbox"}}
            <input type="checkbox" name="{{.Model}}" {{if or .Value .Checked}}checked{{end}} {{if .Readonly}}disabled{{end}}> {{.Label}}
          {{end}}
          {{if eq .Type "radio"}}
            <div name="radio-{{$fieldModel}}">
            <label></label>
            {{range .Options}}
            <label><input type="radio" name="{{$fieldModel}}" value="{{.Key}}" {{if .Selected}}checked="checked"{{end}} {{if .Readonly}}disabled{{end}}> {{.Display}}</label>
            {{end}}
            </div>
          {{end}}
          {{if eq .Type "button"}}
            <input type="button" class="button-outline" value="{{.Label}}" name="{{.Model}}">
            <!-- <button class="button-primary" name="{{.Model}}">{{.Label}}</button>  -->
          {{end}}
        </div>
        {{end}}
      </div>
      {{end}}
    </fieldset>
    <div class="row no-print" name="form-button-bar">
      <div class="column">
        <div class="button-bar">
          {{if .SaveCB}}
          <input type="button" class="button-outline md-close" value="Cancel">
          <button class="button-primary md-save">Save</button> 
          {{else}}
          <input type="button" class="button-outline md-close" value="Close">
          {{end}}
        </div>
      </div>
    </div>
  </form>

</div>
<div id="action-grid" class="no-print"></div>

{{if .DeleteCB}}
<div id="confirm-delete" class="md-modal md-effect-1 modal-container">
  <div class="grid-form md-content">
    Delete This Record ?
  </div>
  <div class="row">
    <input type="button" class="button-outline md-close-del column" value="Cancel">
    <button class="button-primary md-confirm-del column">Yes - Delete</button>
  </div>
</div>
<div class="md-overlay-red"></div>
{{end}}`); err != nil {
		panic(err)
	}

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
	<div class="action__item" url="/blogs">
		<div class="action__title">Blogs</div>
		<div class="action__icon"><i class="fa fa-hashtag fa-lg"></i></div>
		<div class="action__text">
			Edit Blog posts.
		</div>
	</div>
	<div class="action__item" url="/categories">
		<div class="action__title">Categories</div>
		<div class="action__icon"><i class="fa fa-cubes fa-lg"></i></div>
		<div class="action__text">
			Setup Product Categories.
		</div>
	</div>
	<div class="action__item" url="/products">
		<div class="action__title">Products</div>
		<div class="action__icon"><i class="fa fa-cube fa-lg"></i></div>
		<div class="action__text">
			Setup Product Details.
		</div>
	</div>
	<div class="action__item" url="/shippings">
		<div class="action__title">Shipping</div>
		<div class="action__icon"><i class="fa fa-truck fa-lg"></i></div>
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
	<div class="action__item" url="/newsletters">
		<div class="action__title">Newsletter</div>
		<div class="action__icon"><i class="fa fa-newspaper-o fa-lg"></i></div>
		<div class="action__text">
			Review and Generate Mailouts.
		</div>
	</div>
	<div class="action__item" url="/referrers">
		<div class="action__title">Referrers</div>
		<div class="action__icon"><i class="fa fa-cloud-download fa-lg"></i></div>
		<div class="action__text">
			Incoming Traffic from Referring Sites.
		</div>
	</div>
	<div class="action__item" url="/linkouts">
		<div class="action__title">Links</div>
		<div class="action__icon"><i class="fa fa-cloud-upload fa-lg"></i></div>
		<div class="action__text">
			Outgoing Traffic to External Sites.
		</div>
	</div>
	<div class="action__item" url="/customers">
		<div class="action__title">Customers</div>
		<div class="action__icon"><i class="fa fa-user fa-lg"></i></div>
		<div class="action__text">
			Customer Database.
		</div>
	</div>
	<div class="action__item" url="/admins">
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
