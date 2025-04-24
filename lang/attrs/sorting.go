package attrs

var SelfClosingTags = []string{
	"area", "base", "br", "col",
	"embed", "hr", "img", "input",
	"link", "meta", "param", "source",
	"track", "wbr", "path", "rect",
	"polygon", "stop", "ellipse",
}

var AttrSortOrder = []string{
	"id", "class", "name",

	"src", "rel", "type", "href", "action", "formaction", "open", "download",

	"width", "height", "cols", "rows", "colspan", "rowspan", "align", "border", "bgcolor",

	"label", "for", "tabindex", "accesskey", "style", "title", "alt",

	"value", "placeholder", "checked", "disabled", "readonly",
	"autocomplete", "autofocus", "novalidate", "form", "enctype", "method",

	"srcdoc", "poster", "controls", "autoplay", "muted", "loop", "preload", "media", "ismap",

	"http-equiv", "accept", "accept-charset", "charset", "color", "cite", "content",
	"contenteditable", "coords", "data", "datetime", "default", "defer", "dir",
	"dirname", "draggable", "enterkeyhint", "headers", "hidden", "high",
	"hreflang", "inert", "inputmode", "kind", "list", "low",
	"max", "maxlength", "min", "multiple", "optimum", "pattern", "popover",
	"popovertarget", "popovertargetaction", "reversed", "sandbox",
	"scope", "selected", "shape", "size", "sizes", "span", "spellcheck", "srclang",
	"srcset", "start", "step", "translate", "usemap", "wrap",

	"onabort", "onafterprint", "onbeforeprint", "onbeforeunload", "onblur",
	"oncanplay", "oncanplaythrough", "onchange", "onclick", "oncontextmenu",
	"oncopy", "oncuechange", "oncut", "ondblclick", "ondrag", "ondragend",
	"ondragenter", "ondragleave", "ondragover", "ondragstart", "ondrop",
	"ondurationchange", "onemptied", "onended", "onerror", "onfocus",
	"onhashchange", "oninput", "oninvalid", "onkeydown", "onkeypress", "onkeyup",
	"onload", "onloadeddata", "onloadedmetadata", "onloadstart", "onmousedown",
	"onmousemove", "onmouseout", "onmouseover", "onmouseup", "onmousewheel",
	"onoffline", "ononline", "onpagehide", "onpageshow", "onpaste", "onpause",
	"onplay", "onplaying", "onpopstate", "onprogress", "onratechange", "onreset",
	"onresize", "onscroll", "onsearch", "onseeked", "onseeking", "onselect",
	"onstalled", "onstorage", "onsubmit", "onsuspend", "ontimeupdate", "ontoggle",
	"onunload", "onvolumechange", "onwaiting", "onwheel",
}
