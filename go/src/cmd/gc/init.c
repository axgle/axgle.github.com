<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/init.c</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/gc/init.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;go.h&#34;

/*
 * a function named init is a special case.
 * it is called by the initialization before
 * main is run. to make it unique within a
 * package and also uncallable, the name,
 * normally &#34;pkg.init&#34;, is altered to &#34;pkg.init·1&#34;.
 */
Node*
renameinit(Node *n)
{
	Sym *s;
	static int initgen;

	s = n-&gt;sym;
	if(s == S)
		return n;
	if(strcmp(s-&gt;name, &#34;init&#34;) != 0)
		return n;

	snprint(namebuf, sizeof(namebuf), &#34;init·%d&#34;, ++initgen);
	s = lookup(namebuf);
	return newname(s);
}

/*
 * hand-craft the following initialization code
 *	var initdone·&lt;file&gt; uint8 			(1)
 *	func	Init·&lt;file&gt;()				(2)
 *		if initdone·&lt;file&gt; != 0 {		(3)
 *			if initdone·&lt;file&gt; == 2		(4)
 *				return
 *			throw();			(5)
 *		}
 *		initdone.&lt;file&gt; += 1;			(6)
 *		// over all matching imported symbols
 *			&lt;pkg&gt;.init·&lt;file&gt;()		(7)
 *		{ &lt;init stmts&gt; }			(8)
 *		init·&lt;file&gt;()	// if any		(9)
 *		initdone.&lt;file&gt; += 1;			(10)
 *		return					(11)
 *	}
 * note that this code cannot have an assignment
 * statement or, because of the initflag,  it will
 * be converted into a data statement.
 */
int
anyinit(NodeList *n)
{
	uint32 h;
	Sym *s;
	NodeList *l;

	// are there any interesting init statements
	for(l=n; l; l=l-&gt;next) {
		switch(l-&gt;n-&gt;op) {
		case ODCLFUNC:
		case ODCLCONST:
		case ODCLTYPE:
		case OEMPTY:
			break;
		default:
			return 1;
		}
	}

	// is this main
	if(strcmp(package, &#34;main&#34;) == 0)
		return 1;

	// is there an explicit init function
	snprint(namebuf, sizeof(namebuf), &#34;init·1&#34;);
	s = lookup(namebuf);
	if(s-&gt;def != N)
		return 1;

	// are there any imported init functions
	for(h=0; h&lt;NHASH; h++)
	for(s = hash[h]; s != S; s = s-&gt;link) {
		if(s-&gt;name[0] != &#39;I&#39; || strncmp(s-&gt;name, &#34;Init·&#34;, 6) != 0)
			continue;
		if(s-&gt;def == N)
			continue;
		return 1;
	}

	// then none
	return 0;
}

void
fninit(NodeList *n)
{
	int i;
	Node *gatevar;
	Node *a, *b, *fn;
	NodeList *r;
	uint32 h;
	Sym *s, *initsym;

	if(strcmp(package, &#34;PACKAGE&#34;) == 0) {
		// sys.go or unsafe.go during compiler build
		return;
	}

	n = initfix(n);
	if(!anyinit(n))
		return;

	r = nil;

	// (1)
	snprint(namebuf, sizeof(namebuf), &#34;initdone·&#34;);
	gatevar = newname(lookup(namebuf));
	addvar(gatevar, types[TUINT8], PEXTERN);

	// (2)

	maxarg = 0;
	snprint(namebuf, sizeof(namebuf), &#34;Init·&#34;);

	// this is a botch since we need a known name to
	// call the top level init function out of rt0
	if(strcmp(package, &#34;main&#34;) == 0)
		snprint(namebuf, sizeof(namebuf), &#34;init&#34;);

	fn = nod(ODCLFUNC, N, N);
	initsym = lookup(namebuf);
	fn-&gt;nname = newname(initsym);
	fn-&gt;nname-&gt;ntype = nod(OTFUNC, N, N);
	funchdr(fn);

	// (3)
	a = nod(OIF, N, N);
	a-&gt;ntest = nod(ONE, gatevar, nodintconst(0));
	r = list(r, a);

	// (4)
	b = nod(OIF, N, N);
	b-&gt;ntest = nod(OEQ, gatevar, nodintconst(2));
	b-&gt;nbody = list1(nod(ORETURN, N, N));
	a-&gt;nbody = list1(b);

	// (5)
	b = syslook(&#34;throwinit&#34;, 0);
	b = nod(OCALL, b, N);
	a-&gt;nbody = list(a-&gt;nbody, b);

	// (6)
	a = nod(OAS, gatevar, nodintconst(1));
	r = list(r, a);

	// (7)
	for(h=0; h&lt;NHASH; h++)
	for(s = hash[h]; s != S; s = s-&gt;link) {
		if(s-&gt;name[0] != &#39;I&#39; || strncmp(s-&gt;name, &#34;Init·&#34;, 6) != 0)
			continue;
		if(s-&gt;def == N)
			continue;
		if(s == initsym)
			continue;

		// could check that it is fn of no args/returns
		a = nod(OCALL, s-&gt;def, N);
		r = list(r, a);
	}

	// (8)
	r = concat(r, n);

	// (9)
	// could check that it is fn of no args/returns
	for(i=1;; i++) {
		snprint(namebuf, sizeof(namebuf), &#34;init·%d&#34;, i);
		s = lookup(namebuf);
		if(s-&gt;def == N)
			break;
		a = nod(OCALL, s-&gt;def, N);
		r = list(r, a);
	}

	// (10)
	a = nod(OAS, gatevar, nodintconst(2));
	r = list(r, a);

	// (11)
	a = nod(ORETURN, N, N);
	r = list(r, a);

	exportsym(fn-&gt;nname);

	fn-&gt;nbody = r;
	funcbody(fn);
	typecheck(&amp;fn, Etop);
	funccompile(fn);
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
