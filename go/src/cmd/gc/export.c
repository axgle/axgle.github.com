<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/export.c</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/gc/export.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include	&#34;go.h&#34;
#include	&#34;y.tab.h&#34;

void dumpsym(Sym*);

void
addexportsym(Node *n)
{
	exportlist = list(exportlist, n);
}

void
exportsym(Node *n)
{
	if(n == N || n-&gt;sym == S)
		return;
	if(n-&gt;sym-&gt;flags &amp; (SymExport|SymPackage)) {
		if(n-&gt;sym-&gt;flags &amp; SymPackage)
			yyerror(&#34;export/package mismatch: %S&#34;, n-&gt;sym);
		return;
	}
	n-&gt;sym-&gt;flags |= SymExport;

	addexportsym(n);
}

void
packagesym(Node *n)
{
	if(n == N || n-&gt;sym == S)
		return;
	if(n-&gt;sym-&gt;flags &amp; (SymExport|SymPackage)) {
		if(n-&gt;sym-&gt;flags &amp; SymExport)
			yyerror(&#34;export/package mismatch: %S&#34;, n-&gt;sym);
		return;
	}
	n-&gt;sym-&gt;flags |= SymPackage;

	addexportsym(n);
}

int
exportname(char *s)
{
	Rune r;

	if((uchar)s[0] &lt; Runeself)
		return &#39;A&#39; &lt;= s[0] &amp;&amp; s[0] &lt;= &#39;Z&#39;;
	chartorune(&amp;r, s);
	return isupperrune(r);
}

void
autoexport(Node *n, int ctxt)
{
	if(n == N || n-&gt;sym == S)
		return;
	if((ctxt != PEXTERN &amp;&amp; ctxt != PFUNC) || dclcontext != PEXTERN)
		return;
	if(n-&gt;ntype &amp;&amp; n-&gt;ntype-&gt;op == OTFUNC &amp;&amp; n-&gt;ntype-&gt;left)	// method
		return;
	if(exportname(n-&gt;sym-&gt;name) || strcmp(n-&gt;sym-&gt;name, &#34;init&#34;) == 0)
		exportsym(n);
	else
		packagesym(n);
}

void
dumpprereq(Type *t)
{
	if(t == T)
		return;

	if(t-&gt;printed || t == types[t-&gt;etype])
		return;
	t-&gt;printed = 1;

	if(t-&gt;sym != S &amp;&amp; t-&gt;etype != TFIELD)
		dumpsym(t-&gt;sym);
	dumpprereq(t-&gt;type);
	dumpprereq(t-&gt;down);
}

void
dumpexportconst(Sym *s)
{
	Node *n;
	Type *t;

	n = s-&gt;def;
	typecheck(&amp;n, Erv);
	if(n == N || n-&gt;op != OLITERAL)
		fatal(&#34;dumpexportconst: oconst nil: %S&#34;, s);

	t = n-&gt;type;	// may or may not be specified
	if(t != T)
		dumpprereq(t);

	Bprint(bout, &#34;\t&#34;);
	Bprint(bout, &#34;const %lS&#34;, s);
	if(t != T &amp;&amp; !isideal(t))
		Bprint(bout, &#34; %#T&#34;, t);
	Bprint(bout, &#34; = &#34;);

	switch(n-&gt;val.ctype) {
	default:
		fatal(&#34;dumpexportconst: unknown ctype: %S&#34;, s);
	case CTINT:
		Bprint(bout, &#34;%B\n&#34;, n-&gt;val.u.xval);
		break;
	case CTBOOL:
		if(n-&gt;val.u.bval)
			Bprint(bout, &#34;true\n&#34;);
		else
			Bprint(bout, &#34;false\n&#34;);
		break;
	case CTFLT:
		Bprint(bout, &#34;%F\n&#34;, n-&gt;val.u.fval);
		break;
	case CTSTR:
		Bprint(bout, &#34;\&#34;%Z\&#34;\n&#34;, n-&gt;val.u.sval);
		break;
	}
}

void
dumpexportvar(Sym *s)
{
	Node *n;
	Type *t;

	n = s-&gt;def;
	typecheck(&amp;n, Erv);
	if(n == N || n-&gt;type == T) {
		yyerror(&#34;variable exported but not defined: %S&#34;, s);
		return;
	}

	t = n-&gt;type;
	dumpprereq(t);

	Bprint(bout, &#34;\t&#34;);
	if(t-&gt;etype == TFUNC &amp;&amp; n-&gt;class == PFUNC)
		Bprint(bout, &#34;func %lS %#hhT&#34;, s, t);
	else
		Bprint(bout, &#34;var %lS %#T&#34;, s, t);
	Bprint(bout, &#34;\n&#34;);
}

void
dumpexporttype(Sym *s)
{
	Type *t;

	t = s-&gt;def-&gt;type;
	dumpprereq(t);
	Bprint(bout, &#34;\t&#34;);
	switch (t-&gt;etype) {
	case TFORW:
		yyerror(&#34;export of incomplete type %T&#34;, t);
		return;
	}
	Bprint(bout, &#34;type %#T %l#T\n&#34;,  t, t);
}

void
dumpsym(Sym *s)
{
	Type *f, *t;

	if(s-&gt;flags &amp; SymExported)
		return;
	s-&gt;flags |= SymExported;

	if(s-&gt;def == N) {
		yyerror(&#34;unknown export symbol: %S&#34;, s);
		return;
	}
	switch(s-&gt;def-&gt;op) {
	default:
		yyerror(&#34;unexpected export symbol: %O %S&#34;, s-&gt;def-&gt;op, s);
		break;
	case OLITERAL:
		dumpexportconst(s);
		break;
	case OTYPE:
		t = s-&gt;def-&gt;type;
		// TODO(rsc): sort methods by name
		for(f=t-&gt;method; f!=T; f=f-&gt;down)
			dumpprereq(f);

		dumpexporttype(s);
		for(f=t-&gt;method; f!=T; f=f-&gt;down)
			Bprint(bout, &#34;\tfunc (%#T) %hS %#hhT\n&#34;,
				f-&gt;type-&gt;type-&gt;type, f-&gt;sym, f-&gt;type);
		break;
	case ONAME:
		dumpexportvar(s);
		break;
	}
}

void
dumptype(Type *t)
{
	// no need to re-dump type if already exported
	if(t-&gt;printed)
		return;

	// no need to dump type if it&#39;s not ours (was imported)
	if(t-&gt;sym != S &amp;&amp; t-&gt;sym-&gt;def == typenod(t) &amp;&amp; !t-&gt;local)
		return;

	Bprint(bout, &#34;type %#T %l#T\n&#34;,  t, t);
}

void
dumpexport(void)
{
	NodeList *l;
	int32 lno;

	lno = lineno;

	Bprint(bout, &#34;   import\n&#34;);
	Bprint(bout, &#34;\n$$  // exports\n&#34;);

	Bprint(bout, &#34;    package %s\n&#34;, package);

	for(l=exportlist; l; l=l-&gt;next) {
		lineno = l-&gt;n-&gt;lineno;
		dumpsym(l-&gt;n-&gt;sym);
	}

	Bprint(bout, &#34;\n$$  // local types\n&#34;);

	for(l=typelist; l; l=l-&gt;next) {
		lineno = l-&gt;n-&gt;lineno;
		dumptype(l-&gt;n-&gt;type);
	}

	Bprint(bout, &#34;\n$$\n&#34;);

	lineno = lno;
}

/*
 * import
 */

/*
 * return the sym for ss, which should match lexical
 */
Sym*
importsym(Sym *s, int op)
{
	if(s-&gt;def != N &amp;&amp; s-&gt;def-&gt;op != op)
		redeclare(s, &#34;during import&#34;);

	// mark the symbol so it is not reexported
	if(s-&gt;def == N) {
		if(exportname(s-&gt;name))
			s-&gt;flags |= SymExport;
		else
			s-&gt;flags |= SymPackage;	// package scope
		s-&gt;flags |= SymImported;
	}
	return s;
}

/*
 * return the type pkg.name, forward declaring if needed
 */
Type*
pkgtype(Sym *s)
{
	Type *t;

	importsym(s, OTYPE);
	if(s-&gt;def == N || s-&gt;def-&gt;op != OTYPE) {
		t = typ(TFORW);
		t-&gt;sym = s;
		s-&gt;def = typenod(t);
	}
	if(s-&gt;def-&gt;type == T)
		yyerror(&#34;pkgtype %lS&#34;, s);
	return s-&gt;def-&gt;type;
}

static int
mypackage(Sym *s)
{
	// we import all definitions for runtime.
	// lowercase ones can only be used by the compiler.
	return strcmp(s-&gt;package, package) == 0
		|| strcmp(s-&gt;package, &#34;runtime&#34;) == 0;
}

void
importconst(Sym *s, Type *t, Node *n)
{
	Node *n1;

	if(!exportname(s-&gt;name) &amp;&amp; !mypackage(s))
		return;
	importsym(s, OLITERAL);
	convlit(&amp;n, t);
	if(s-&gt;def != N) {
		// TODO: check if already the same.
		return;
	}

	if(n-&gt;op != OLITERAL) {
		yyerror(&#34;expression must be a constant&#34;);
		return;
	}
	if(n-&gt;sym != S) {
		n1 = nod(OXXX, N, N);
		*n1 = *n;
		n = n1;
	}
	n-&gt;sym = s;
	declare(n, PEXTERN);

	if(debug[&#39;E&#39;])
		print(&#34;import const %S\n&#34;, s);
}

void
importvar(Sym *s, Type *t, int ctxt)
{
	Node *n;

	if(!exportname(s-&gt;name) &amp;&amp; !mypackage(s))
		return;

	importsym(s, ONAME);
	if(s-&gt;def != N &amp;&amp; s-&gt;def-&gt;op == ONAME) {
		if(cvttype(t, s-&gt;def-&gt;type))
			return;
		yyerror(&#34;inconsistent definition for var %S during import\n\t%T\n\t%T&#34;,
			s, s-&gt;def-&gt;type, t);
	}
	n = newname(s);
	n-&gt;type = t;
	declare(n, ctxt);

	if(debug[&#39;E&#39;])
		print(&#34;import var %S %lT\n&#34;, s, t);
}

void
importtype(Type *pt, Type *t)
{
	if(pt != T &amp;&amp; t != T)
		typedcl2(pt, t);

	if(debug[&#39;E&#39;])
		print(&#34;import type %T %lT\n&#34;, pt, t);
}

void
importmethod(Sym *s, Type *t)
{
	checkwidth(t);
	addmethod(s, t, 0);
}

/*
 * ******* import *******
 */

void
checkimports(void)
{
	Sym *s;
	Type *t, *t1;
	uint32 h;
	int et;

return;

	for(h=0; h&lt;NHASH; h++)
	for(s = hash[h]; s != S; s = s-&gt;link) {
		if(s-&gt;def == N || s-&gt;def-&gt;op != OTYPE)
			continue;
		t = s-&gt;def-&gt;type;
		if(t == T)
			continue;

		et = t-&gt;etype;
		switch(t-&gt;etype) {
		case TFORW:
			print(&#34;ci-1: %S %lT\n&#34;, s, t);
			break;

		case TPTR32:
		case TPTR64:
			if(t-&gt;type == T) {
				print(&#34;ci-2: %S %lT\n&#34;, s, t);
				break;
			}

			t1 = t-&gt;type;
			if(t1 == T) {
				print(&#34;ci-3: %S %lT\n&#34;, s, t1);
				break;
			}

			et = t1-&gt;etype;
			if(et == TFORW) {
				print(&#34;%L: ci-4: %S %lT\n&#34;, lineno, s, t);
				break;
			}
			break;
		}
	}
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
