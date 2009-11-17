<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/closure.c</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

</head>

<body>

  <script>
    // Catch 'enter' key down events and trigger the search form submission.
    function codesearchKeyDown(event) {
      if (event.which == 13) {
        var form = document.getElementById('codesearch');
        var query = document.getElementById('codesearchQuery');
        form.q.value = "lang:go package:go.googlecode.com " + query.value;
        document.getElementById('codesearch').submit();
}      return true;
}
    // Capture the submission event and construct the query parameter.
    function codeSearchSubmit() {
      var query = document.getElementById('codesearchQuery');
      var form = document.getElementById('codesearch');
      form.q.value = "lang:go package:go.googlecode.com " + query.value;
      return true;
}  </script>

<div id="topnav">
  <table summary="">
    <tr>
      <td id="headerImage">
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
      </td>
      <td>
        <div id="headerDocSetTitle">The Go Programming Language</div>
      </td>
      <td>
        <!-- <table>
          <tr>
            <td>
              <! The input box is outside of the form because we want to add
              a couple of restricts to the query before submitting. If we just
              add the restricts to the text box before submitting, then they
              appear in the box when the user presses 'back'. Thus we use a
              hidden field in the form. However, there's no way to stop the
              non-hidden text box from also submitting a value unless we move
              it outside of the form
              <input type="search" id="codesearchQuery" value="" size="30" onkeydown="return codesearchKeyDown(event);"/>
              <form method="GET" action="http://www.google.com/codesearch" id="codesearch" class="search" onsubmit="return codeSearchSubmit();" style="display:inline;">
                <input type="hidden" name="q" value=""/>
                <input type="submit" value="Code search" />
                <span style="color: red">(TODO: remove for now?)</span>
              </form>
            </td>
          </tr>
          <tr>
            <td>
              <span style="color: gray;">(e.g. &ldquo;pem&rdquo; or &ldquo;xml&rdquo;)</span>
            </td>
          </tr>
        </table> -->
      </td>
    </tr>
  </table>
</div>

<div id="linkList">
  <ul>
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Help</li>
    <li>#go-nuts on irc.freenode.net</li>
    <li><a href="http://groups.google.com/group/golang-nuts">Go Nuts mailing list</a></li>
    <li><a href="http://code.google.com/p/go/issues/list">Issue tracker</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Go code search</li>
    <form method="GET" action="http://golang.org/search" class="search">
    <input type="search" name="q" value="" size="25" style="width:80%; max-width:200px" />
    <input type="submit" value="Go" />
    </form>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Last update</li>
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/gc/closure.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * function literals aka closures
 */

#include &#34;go.h&#34;

void
closurehdr(Node *ntype)
{
	Node *n, *name;
	NodeList *l;

	n = nod(OCLOSURE, N, N);
	n-&gt;ntype = ntype;
	n-&gt;funcdepth = funcdepth;

	funchdr(n);

	// steal ntype&#39;s argument names and
	// leave a fresh copy in their place.
	// references to these variables need to
	// refer to the variables in the external
	// function declared below; see walkclosure.
	n-&gt;list = ntype-&gt;list;
	n-&gt;rlist = ntype-&gt;rlist;
	ntype-&gt;list = nil;
	ntype-&gt;rlist = nil;
	for(l=n-&gt;list; l; l=l-&gt;next) {
		name = l-&gt;n-&gt;left;
		if(name)
			name = newname(name-&gt;sym);
		ntype-&gt;list = list(ntype-&gt;list, nod(ODCLFIELD, name, l-&gt;n-&gt;right));
	}
	for(l=n-&gt;rlist; l; l=l-&gt;next) {
		name = l-&gt;n-&gt;left;
		if(name)
			name = newname(name-&gt;sym);
		ntype-&gt;rlist = list(ntype-&gt;rlist, nod(ODCLFIELD, name, l-&gt;n-&gt;right));
	}
}

Node*
closurebody(NodeList *body)
{
	Node *func, *v;
	NodeList *l;

	if(body == nil)
		body = list1(nod(OEMPTY, N, N));

	func = curfn;
	l = func-&gt;dcl;
	func-&gt;nbody = body;
	funcbody(func);

	// closure-specific variables are hanging off the
	// ordinary ones in the symbol table; see oldname.
	// unhook them.
	// make the list of pointers for the closure call.
	for(l=func-&gt;cvars; l; l=l-&gt;next) {
		v = l-&gt;n;
		v-&gt;closure-&gt;closure = v-&gt;outer;
		v-&gt;heapaddr = nod(OADDR, oldname(v-&gt;sym), N);
	}

	return func;
}

void
typecheckclosure(Node *func)
{
	Node *oldfn;
	NodeList *l;
	Node *v;

	oldfn = curfn;
	typecheck(&amp;func-&gt;ntype, Etype);
	func-&gt;type = func-&gt;ntype-&gt;type;
	if(func-&gt;type != T) {
		curfn = func;
		typechecklist(func-&gt;nbody, Etop);
		curfn = oldfn;
	}

	// type check the &amp; of closed variables outside the closure,
	// so that the outer frame also grabs them and knows they
	// escape.
	func-&gt;enter = nil;
	for(l=func-&gt;cvars; l; l=l-&gt;next) {
		v = l-&gt;n;
		if(v-&gt;type == T) {
			// if v-&gt;type is nil, it means v looked like it was
			// going to be used in the closure but wasn&#39;t.
			// this happens because when parsing a, b, c := f()
			// the a, b, c gets parsed as references to older
			// a, b, c before the parser figures out this is a
			// declaration.
			v-&gt;op = 0;
			continue;
		}
		typecheck(&amp;v-&gt;heapaddr, Erv);
		func-&gt;enter = list(func-&gt;enter, v-&gt;heapaddr);
		v-&gt;heapaddr = N;
	}
}

Node*
walkclosure(Node *func, NodeList **init)
{
	int narg;
	Node *xtype, *v, *addr, *xfunc, *call, *clos;
	NodeList *l, *in;
	static int closgen;

	/*
	 * wrap body in external function
	 * with extra closure parameters.
	 */
	xtype = nod(OTFUNC, N, N);

	// each closure variable has a corresponding
	// address parameter.
	narg = 0;
	for(l=func-&gt;cvars; l; l=l-&gt;next) {
		v = l-&gt;n;
		if(v-&gt;op == 0)
			continue;
		addr = nod(ONAME, N, N);
		snprint(namebuf, sizeof namebuf, &#34;&amp;%s&#34;, v-&gt;sym-&gt;name);
		addr-&gt;sym = lookup(namebuf);
		addr-&gt;ntype = nod(OIND, typenod(v-&gt;type), N);
		addr-&gt;class = PPARAM;
		addr-&gt;addable = 1;
		addr-&gt;ullman = 1;
		narg++;

		v-&gt;heapaddr = addr;

		xtype-&gt;list = list(xtype-&gt;list, nod(ODCLFIELD, addr, addr-&gt;ntype));
	}

	// then a dummy arg where the closure&#39;s caller pc sits
	xtype-&gt;list = list(xtype-&gt;list, nod(ODCLFIELD, N, typenod(types[TUINTPTR])));

	// then the function arguments
	xtype-&gt;list = concat(xtype-&gt;list, func-&gt;list);
	xtype-&gt;rlist = concat(xtype-&gt;rlist, func-&gt;rlist);

	// create the function
	xfunc = nod(ODCLFUNC, N, N);
	snprint(namebuf, sizeof namebuf, &#34;_f%.3ld&#34;, ++closgen);
	xfunc-&gt;nname = newname(lookup(namebuf));
	xfunc-&gt;nname-&gt;ntype = xtype;
	declare(xfunc-&gt;nname, PFUNC);
	xfunc-&gt;nname-&gt;funcdepth = func-&gt;funcdepth;
	xfunc-&gt;funcdepth = func-&gt;funcdepth;
	xfunc-&gt;nbody = func-&gt;nbody;
	xfunc-&gt;dcl = func-&gt;dcl;
	if(xfunc-&gt;nbody == nil)
		fatal(&#34;empty body - won&#39;t generate any code&#34;);
	typecheck(&amp;xfunc, Etop);
	closures = list(closures, xfunc);

	// prepare call of sys.closure that turns external func into func literal value.
	clos = syslook(&#34;closure&#34;, 1);
	clos-&gt;type = T;
	clos-&gt;ntype = nod(OTFUNC, N, N);
	in = list1(nod(ODCLFIELD, N, typenod(types[TINT])));	// siz
	in = list(in, nod(ODCLFIELD, N, xtype));
	for(l=func-&gt;cvars; l; l=l-&gt;next) {
		if(l-&gt;n-&gt;op == 0)
			continue;
		in = list(in, nod(ODCLFIELD, N, l-&gt;n-&gt;heapaddr-&gt;ntype));
	}
	clos-&gt;ntype-&gt;list = in;
	clos-&gt;ntype-&gt;rlist = list1(nod(ODCLFIELD, N, typenod(func-&gt;type)));
	typecheck(&amp;clos, Erv);

	call = nod(OCALL, clos, N);
	if(narg*widthptr &gt; 100)
		yyerror(&#34;closure needs too many variables; runtime will reject it&#34;);
	in = list1(nodintconst(narg*widthptr));
	in = list(in, xfunc-&gt;nname);
	in = concat(in, func-&gt;enter);
	call-&gt;list = in;

	typecheck(&amp;call, Erv);
	walkexpr(&amp;call, init);
	return call;
}
</pre>

</div>

<div id="footer">
<p>Except as noted, this content is
   licensed under <a href="http://creativecommons.org/licenses/by/3.0/">
   Creative Commons Attribution 3.0</a>.
</div>

<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-11222381-2");
pageTracker._trackPageview();
</script>
</body>
</html>
<!-- generated at Thu Nov 12 15:42:51 PST 2009 -->
