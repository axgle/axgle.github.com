<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/subr.c</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/gc/subr.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include	&#34;go.h&#34;
#include	&#34;md5.h&#34;
#include	&#34;y.tab.h&#34;
#include	&#34;opnames.h&#34;

typedef struct Error Error;
struct Error
{
	int lineno;
	int seq;
	char *msg;
};
static Error *err;
static int nerr;
static int merr;

void
errorexit(void)
{
	flusherrors();
	if(outfile)
		remove(outfile);
	exit(1);
}

extern int yychar;
int
parserline(void)
{
	if(yychar != 0 &amp;&amp; yychar != -2)	// parser has one symbol lookahead
		return prevlineno;
	return lineno;
}

static void
adderr(int line, char *fmt, va_list arg)
{
	Fmt f;
	Error *p;

	fmtstrinit(&amp;f);
	fmtprint(&amp;f, &#34;%L: &#34;, line);
	fmtvprint(&amp;f, fmt, arg);
	fmtprint(&amp;f, &#34;\n&#34;);

	if(nerr &gt;= merr) {
		if(merr == 0)
			merr = 16;
		else
			merr *= 2;
		p = realloc(err, merr*sizeof err[0]);
		if(p == nil) {
			merr = nerr;
			flusherrors();
			print(&#34;out of memory\n&#34;);
			errorexit();
		}
		err = p;
	}
	err[nerr].seq = nerr;
	err[nerr].lineno = line;
	err[nerr].msg = fmtstrflush(&amp;f);
	nerr++;
}

static int
errcmp(const void *va, const void *vb)
{
	Error *a, *b;

	a = (Error*)va;
	b = (Error*)vb;
	if(a-&gt;lineno != b-&gt;lineno)
		return a-&gt;lineno - b-&gt;lineno;
	if(a-&gt;seq != b-&gt;seq)
		return a-&gt;seq - b-&gt;seq;
	return 0;
}

void
flusherrors(void)
{
	int i;

	if(nerr == 0)
		return;
	qsort(err, nerr, sizeof err[0], errcmp);
	for(i=0; i&lt;nerr; i++)
		print(&#34;%s&#34;, err[i].msg);
	nerr = 0;
}

static void
hcrash(void)
{
	if(debug[&#39;h&#39;]) {
		flusherrors();
		if(outfile)
			unlink(outfile);
		*(int*)0 = 0;
	}
}

void
yyerrorl(int line, char *fmt, ...)
{
	va_list arg;

	va_start(arg, fmt);
	adderr(line, fmt, arg);
	va_end(arg);

	hcrash();
	nerrors++;
	if(nerrors &gt;= 10 &amp;&amp; !debug[&#39;e&#39;])
		fatal(&#34;too many errors&#34;);
}

void
yyerror(char *fmt, ...)
{
	va_list arg;

	if(strcmp(fmt, &#34;syntax error&#34;) == 0) {
		yyerrorl(lexlineno, &#34;syntax error near %s&#34;, lexbuf);
		nsyntaxerrors++;
		return;
	}

	va_start(arg, fmt);
	adderr(parserline(), fmt, arg);
	va_end(arg);

	hcrash();
	nerrors++;
	if(nerrors &gt;= 10 &amp;&amp; !debug[&#39;e&#39;])
		fatal(&#34;too many errors&#34;);
}

void
warn(char *fmt, ...)
{
	va_list arg;

	va_start(arg, fmt);
	adderr(parserline(), fmt, arg);
	va_end(arg);

	hcrash();
}

void
fatal(char *fmt, ...)
{
	va_list arg;

	flusherrors();

	print(&#34;%L: fatal error: &#34;, lineno);
	va_start(arg, fmt);
	vfprint(1, fmt, arg);
	va_end(arg);
	print(&#34;\n&#34;);

	hcrash();
	errorexit();
}

void
linehist(char *file, int32 off, int relative)
{
	Hist *h;
	char *cp;

	if(debug[&#39;i&#39;]) {
		if(file != nil) {
			if(off &lt; 0)
				print(&#34;pragma %s at line %L\n&#34;, file, lineno);
			else
				print(&#34;import %s at line %L\n&#34;, file, lineno);
		} else
			print(&#34;end of import at line %L\n&#34;, lineno);
	}

	if(off &lt; 0 &amp;&amp; file[0] != &#39;/&#39; &amp;&amp; !relative) {
		cp = mal(strlen(file) + strlen(pathname) + 2);
		sprint(cp, &#34;%s/%s&#34;, pathname, file);
		file = cp;
	}

	h = mal(sizeof(Hist));
	h-&gt;name = file;
	h-&gt;line = lexlineno;
	h-&gt;offset = off;
	h-&gt;link = H;
	if(ehist == H) {
		hist = h;
		ehist = h;
		return;
	}
	ehist-&gt;link = h;
	ehist = h;
}

int32
setlineno(Node *n)
{
	int32 lno;

	lno = lineno;
	if(n != N)
	switch(n-&gt;op) {
	case ONAME:
	case OTYPE:
	case OPACK:
	case OLITERAL:
	case ONONAME:
		break;
	default:
		lineno = n-&gt;lineno;
		if(lineno == 0) {
			if(debug[&#39;K&#39;])
				warn(&#34;setlineno: line 0&#34;);
			lineno = lno;
		}
	}
	return lno;
}

uint32
stringhash(char *p)
{
	int32 h;
	int c;

	h = 0;
	for(;;) {
		c = *p++;
		if(c == 0)
			break;
		h = h*PRIME1 + c;
	}

	if(h &lt; 0) {
		h = -h;
		if(h &lt; 0)
			h = 0;
	}
	return h;
}

Sym*
lookup(char *p)
{
	Sym *s;
	uint32 h;
	int c;

	h = stringhash(p) % NHASH;
	c = p[0];

	for(s = hash[h]; s != S; s = s-&gt;link) {
		if(s-&gt;name[0] != c)
			continue;
		if(strcmp(s-&gt;name, p) == 0)
			if(s-&gt;package &amp;&amp; strcmp(s-&gt;package, package) == 0)
				return s;
	}

	s = mal(sizeof(*s));
	s-&gt;name = mal(strlen(p)+1);
	s-&gt;package = package;
	s-&gt;lexical = LNAME;

	strcpy(s-&gt;name, p);

	s-&gt;link = hash[h];
	hash[h] = s;

	return s;
}

Sym*
pkglookup(char *name, char *pkg)
{
	Sym *s;
	uint32 h;
	int c;

	h = stringhash(name) % NHASH;
	c = name[0];
	for(s = hash[h]; s != S; s = s-&gt;link) {
		if(s-&gt;name[0] != c)
			continue;
		if(strcmp(s-&gt;name, name) == 0)
			if(s-&gt;package &amp;&amp; strcmp(s-&gt;package, pkg) == 0)
				return s;
	}

	s = mal(sizeof(*s));
	s-&gt;name = mal(strlen(name)+1);
	strcpy(s-&gt;name, name);

	// botch - should probably try to reuse the pkg string
	s-&gt;package = mal(strlen(pkg)+1);
	strcpy(s-&gt;package, pkg);

	s-&gt;link = hash[h];
	hash[h] = s;
	s-&gt;lexical = LNAME;

	return s;
}

Sym*
restrictlookup(char *name, char *pkg)
{
	if(!exportname(name) &amp;&amp; strcmp(pkg, package) != 0)
		yyerror(&#34;cannot refer to %s.%s&#34;, pkg, name);
	return pkglookup(name, pkg);
}


// find all the exported symbols in package opkg
// and make them available in the current package
void
importdot(Sym *opkg, Node *pack)
{
	Sym *s, *s1;
	uint32 h;
	int c, n;

	if(strcmp(opkg-&gt;name, package) == 0)
		return;

	n = 0;
	c = opkg-&gt;name[0];
	for(h=0; h&lt;NHASH; h++) {
		for(s = hash[h]; s != S; s = s-&gt;link) {
			if(s-&gt;package[0] != c)
				continue;
			if(!exportname(s-&gt;name) || utfrune(s-&gt;name, 0xb7))	// 0xb7 = center dot
				continue;
			if(strcmp(s-&gt;package, opkg-&gt;name) != 0)
				continue;
			s1 = lookup(s-&gt;name);
			if(s1-&gt;def != N) {
				redeclare(s1, &#34;during import&#34;);
				continue;
			}
			s1-&gt;def = s-&gt;def;
			s1-&gt;def-&gt;pack = pack;
			n++;
		}
	}
	if(n == 0) {
		// can&#39;t possibly be used - there were no symbols
		yyerrorl(pack-&gt;lineno, &#34;imported and not used: %s&#34;, pack-&gt;sym-&gt;name);
	}
}

void
gethunk(void)
{
	char *h;
	int32 nh;

	nh = NHUNK;
	if(thunk &gt;= 10L*NHUNK)
		nh = 10L*NHUNK;
	h = (char*)malloc(nh);
	if(h == (char*)-1) {
		flusherrors();
		yyerror(&#34;out of memory&#34;);
		errorexit();
	}
	hunk = h;
	nhunk = nh;
	thunk += nh;
}

void*
mal(int32 n)
{
	void *p;

	while((uintptr)hunk &amp; MAXALIGN) {
		hunk++;
		nhunk--;
	}
	while(nhunk &lt; n)
		gethunk();

	p = hunk;
	nhunk -= n;
	hunk += n;
	memset(p, 0, n);
	return p;
}

void*
remal(void *p, int32 on, int32 n)
{
	void *q;

	q = (uchar*)p + on;
	if(q != hunk || nhunk &lt; n) {
		while(nhunk &lt; on+n)
			gethunk();
		memmove(hunk, p, on);
		p = hunk;
		hunk += on;
		nhunk -= on;
	}
	hunk += n;
	nhunk -= n;
	return p;
}

Node*
nod(int op, Node *nleft, Node *nright)
{
	Node *n;

	n = mal(sizeof(*n));
	n-&gt;op = op;
	n-&gt;left = nleft;
	n-&gt;right = nright;
	n-&gt;lineno = parserline();
	n-&gt;xoffset = BADWIDTH;
	return n;
}

int
algtype(Type *t)
{
	int a;

	if(issimple[t-&gt;etype] || isptr[t-&gt;etype] || t-&gt;etype == TCHAN || t-&gt;etype == TFUNC || t-&gt;etype == TMAP)
		a = AMEM;	// just bytes (int, ptr, etc)
	else if(t-&gt;etype == TSTRING)
		a = ASTRING;	// string
	else if(isnilinter(t))
		a = ANILINTER;	// nil interface
	else if(t-&gt;etype == TINTER)
		a = AINTER;	// interface
	else
		a = ANOEQ;	// just bytes, but no hash/eq
	return a;
}

Type*
maptype(Type *key, Type *val)
{
	Type *t;

	if(key != nil &amp;&amp; key-&gt;etype != TANY &amp;&amp; algtype(key) == ANOEQ) {
		if(key-&gt;etype == TFORW) {
			// map[key] used during definition of key.
			// postpone check until key is fully defined.
			// if there are multiple uses of map[key]
			// before key is fully defined, the error
			// will only be printed for the first one.
			// good enough.
			if(key-&gt;maplineno == 0)
				key-&gt;maplineno = lineno;
		} else
			yyerror(&#34;invalid map key type %T&#34;, key);
	}
	t = typ(TMAP);
	t-&gt;down = key;
	t-&gt;type = val;
	return t;
}

int
iskeytype(Type *t)
{
	return algtype(t) != ANOEQ;
}

Type*
typ(int et)
{
	Type *t;

	t = mal(sizeof(*t));
	t-&gt;etype = et;
	t-&gt;width = BADWIDTH;
	t-&gt;lineno = lineno;
	return t;
}


Type*
sortinter(Type *t)
{
	return t;
}

Node*
nodintconst(int64 v)
{
	Node *c;

	c = nod(OLITERAL, N, N);
	c-&gt;addable = 1;
	c-&gt;val.u.xval = mal(sizeof(*c-&gt;val.u.xval));
	mpmovecfix(c-&gt;val.u.xval, v);
	c-&gt;val.ctype = CTINT;
	c-&gt;type = types[TIDEAL];
	ullmancalc(c);
	return c;
}

void
nodconst(Node *n, Type *t, int64 v)
{
	memset(n, 0, sizeof(*n));
	n-&gt;op = OLITERAL;
	n-&gt;addable = 1;
	ullmancalc(n);
	n-&gt;val.u.xval = mal(sizeof(*n-&gt;val.u.xval));
	mpmovecfix(n-&gt;val.u.xval, v);
	n-&gt;val.ctype = CTINT;
	n-&gt;type = t;

	if(isfloat[t-&gt;etype])
		fatal(&#34;nodconst: bad type %T&#34;, t);
}

Node*
nodnil(void)
{
	Node *c;

	c = nodintconst(0);
	c-&gt;val.ctype = CTNIL;
	c-&gt;type = types[TNIL];
	return c;
}

Node*
nodbool(int b)
{
	Node *c;

	c = nodintconst(0);
	c-&gt;val.ctype = CTBOOL;
	c-&gt;val.u.bval = b;
	c-&gt;type = idealbool;
	return c;
}

Type*
aindex(Node *b, Type *t)
{
	NodeList *init;
	Type *r;
	int bound;

	bound = -1;	// open bound
	init = nil;
	typecheck(&amp;b, Erv);
	if(b != nil) {
		switch(consttype(b)) {
		default:
			yyerror(&#34;array bound must be an integer expression&#34;);
			break;
		case CTINT:
			bound = mpgetfix(b-&gt;val.u.xval);
			if(bound &lt; 0)
				yyerror(&#34;array bound must be non negative&#34;);
			break;
		}
	}

	// fixed array
	r = typ(TARRAY);
	r-&gt;type = t;
	r-&gt;bound = bound;
	return r;
}

void
indent(int dep)
{
	int i;

	for(i=0; i&lt;dep; i++)
		print(&#34;.   &#34;);
}

void
dodumplist(NodeList *l, int dep)
{
	for(; l; l=l-&gt;next)
		dodump(l-&gt;n, dep);
}

void
dodump(Node *n, int dep)
{
	if(n == N)
		return;

	indent(dep);
	if(dep &gt; 10) {
		print(&#34;...\n&#34;);
		return;
	}

	if(n-&gt;ninit != nil) {
		print(&#34;%O-init\n&#34;, n-&gt;op);
		dodumplist(n-&gt;ninit, dep+1);
		indent(dep);
	}

	switch(n-&gt;op) {
	default:
		print(&#34;%N\n&#34;, n);
		dodump(n-&gt;left, dep+1);
		dodump(n-&gt;right, dep+1);
		break;

	case OTYPE:
		print(&#34;%O %S type=%T\n&#34;, n-&gt;op, n-&gt;sym, n-&gt;type);
		if(n-&gt;type == T &amp;&amp; n-&gt;ntype) {
			indent(dep);
			print(&#34;%O-ntype\n&#34;, n-&gt;op);
			dodump(n-&gt;ntype, dep+1);
		}
		break;

	case OIF:
		print(&#34;%O%J\n&#34;, n-&gt;op, n);
		dodump(n-&gt;ntest, dep+1);
		if(n-&gt;nbody != nil) {
			indent(dep);
			print(&#34;%O-then\n&#34;, n-&gt;op);
			dodumplist(n-&gt;nbody, dep+1);
		}
		if(n-&gt;nelse != nil) {
			indent(dep);
			print(&#34;%O-else\n&#34;, n-&gt;op);
			dodumplist(n-&gt;nelse, dep+1);
		}
		break;

	case OSELECT:
		print(&#34;%O%J\n&#34;, n-&gt;op, n);
		dodumplist(n-&gt;nbody, dep+1);
		break;

	case OSWITCH:
	case OFOR:
		print(&#34;%O%J\n&#34;, n-&gt;op, n);
		dodump(n-&gt;ntest, dep+1);

		if(n-&gt;nbody != nil) {
			indent(dep);
			print(&#34;%O-body\n&#34;, n-&gt;op);
			dodumplist(n-&gt;nbody, dep+1);
		}

		if(n-&gt;nincr != N) {
			indent(dep);
			print(&#34;%O-incr\n&#34;, n-&gt;op);
			dodump(n-&gt;nincr, dep+1);
		}
		break;

	case OCASE:
		// the right side points to label of the body
		if(n-&gt;right != N &amp;&amp; n-&gt;right-&gt;op == OGOTO &amp;&amp; n-&gt;right-&gt;left-&gt;op == ONAME)
			print(&#34;%O%J GOTO %N\n&#34;, n-&gt;op, n, n-&gt;right-&gt;left);
		else
			print(&#34;%O%J\n&#34;, n-&gt;op, n);
		dodump(n-&gt;left, dep+1);
		break;

	case OXCASE:
		print(&#34;%N\n&#34;, n);
		dodump(n-&gt;left, dep+1);
		dodump(n-&gt;right, dep+1);
		indent(dep);
		print(&#34;%O-nbody\n&#34;, n-&gt;op);
		dodumplist(n-&gt;nbody, dep+1);
		break;
	}

	if(0 &amp;&amp; n-&gt;ntype != nil) {
		indent(dep);
		print(&#34;%O-ntype\n&#34;, n-&gt;op);
		dodump(n-&gt;ntype, dep+1);
	}
	if(n-&gt;list != nil) {
		indent(dep);
		print(&#34;%O-list\n&#34;, n-&gt;op);
		dodumplist(n-&gt;list, dep+1);
	}
	if(n-&gt;rlist != nil) {
		indent(dep);
		print(&#34;%O-rlist\n&#34;, n-&gt;op);
		dodumplist(n-&gt;rlist, dep+1);
	}
	if(n-&gt;op != OIF &amp;&amp; n-&gt;nbody != nil) {
		indent(dep);
		print(&#34;%O-nbody\n&#34;, n-&gt;op);
		dodumplist(n-&gt;nbody, dep+1);
	}
}

void
dumplist(char *s, NodeList *l)
{
	print(&#34;%s\n&#34;, s);
	dodumplist(l, 1);
}

void
dump(char *s, Node *n)
{
	print(&#34;%s [%p]\n&#34;, s, n);
	dodump(n, 1);
}

static char*
goopnames[] =
{
	[OADDR]		= &#34;&amp;&#34;,
	[OADD]		= &#34;+&#34;,
	[OANDAND]	= &#34;&amp;&amp;&#34;,
	[OANDNOT]	= &#34;&amp;^&#34;,
	[OAND]		= &#34;&amp;&#34;,
	[OAS]		= &#34;=&#34;,
	[OAS2]		= &#34;=&#34;,
	[OBREAK]	= &#34;break&#34;,
	[OCAP]		= &#34;cap&#34;,
	[OCASE]		= &#34;case&#34;,
	[OCLOSED]	= &#34;closed&#34;,
	[OCLOSE]	= &#34;close&#34;,
	[OCOM]		= &#34;^&#34;,
	[OCONTINUE]	= &#34;continue&#34;,
	[ODEC]		= &#34;--&#34;,
	[ODEFER]	= &#34;defer&#34;,
	[ODIV]		= &#34;/&#34;,
	[OEQ]		= &#34;==&#34;,
	[OFALL]		= &#34;fallthrough&#34;,
	[OFOR]		= &#34;for&#34;,
	[OFUNC]		= &#34;func&#34;,
	[OGE]		= &#34;&gt;=&#34;,
	[OGOTO]		= &#34;goto&#34;,
	[OGT]		= &#34;&gt;&#34;,
	[OIF]		= &#34;if&#34;,
	[OINC]		= &#34;++&#34;,
	[OIND]		= &#34;*&#34;,
	[OLEN]		= &#34;len&#34;,
	[OLE]		= &#34;&lt;=&#34;,
	[OLSH]		= &#34;&lt;&lt;&#34;,
	[OLT]		= &#34;&lt;&#34;,
	[OMAKE]		= &#34;make&#34;,
	[OMINUS]	= &#34;-&#34;,
	[OMOD]		= &#34;%&#34;,
	[OMUL]		= &#34;*&#34;,
	[ONEW]		= &#34;new&#34;,
	[ONE]		= &#34;!=&#34;,
	[ONOT]		= &#34;!&#34;,
	[OOROR]		= &#34;||&#34;,
	[OOR]		= &#34;|&#34;,
	[OPANICN]	= &#34;panicln&#34;,
	[OPANIC]	= &#34;panic&#34;,
	[OPLUS]		= &#34;+&#34;,
	[OPRINTN]	= &#34;println&#34;,
	[OPRINT]	= &#34;print&#34;,
	[ORANGE]	= &#34;range&#34;,
	[ORECV]		= &#34;&lt;-&#34;,
	[ORETURN]	= &#34;return&#34;,
	[ORSH]		= &#34;&gt;&gt;&#34;,
	[OSELECT]	= &#34;select&#34;,
	[OSEND]		= &#34;&lt;-&#34;,
	[OSUB]		= &#34;-&#34;,
	[OSWITCH]	= &#34;switch&#34;,
	[OXOR]		= &#34;^&#34;,
};

int
Oconv(Fmt *fp)
{
	char buf[500];
	int o;

	o = va_arg(fp-&gt;args, int);
	if((fp-&gt;flags &amp; FmtSharp) &amp;&amp; o &gt;= 0 &amp;&amp; o &lt; nelem(goopnames) &amp;&amp; goopnames[o] != nil)
		return fmtstrcpy(fp, goopnames[o]);
	if(o &lt; 0 || o &gt;= nelem(opnames) || opnames[o] == nil) {
		snprint(buf, sizeof(buf), &#34;O-%d&#34;, o);
		return fmtstrcpy(fp, buf);
	}
	return fmtstrcpy(fp, opnames[o]);
}

int
Lconv(Fmt *fp)
{
	char str[STRINGSZ], s[STRINGSZ];
	struct
	{
		Hist*	incl;	/* start of this include file */
		int32	idel;	/* delta line number to apply to include */
		Hist*	line;	/* start of this #line directive */
		int32	ldel;	/* delta line number to apply to #line */
	} a[HISTSZ];
	int32 lno, d;
	int i, n;
	Hist *h;

	lno = va_arg(fp-&gt;args, int32);

	n = 0;
	for(h=hist; h!=H; h=h-&gt;link) {
		if(h-&gt;offset &lt; 0)
			continue;
		if(lno &lt; h-&gt;line)
			break;
		if(h-&gt;name) {
			if(n &lt; HISTSZ) {	/* beginning of file */
				a[n].incl = h;
				a[n].idel = h-&gt;line;
				a[n].line = 0;
			}
			n++;
			continue;
		}
		n--;
		if(n &gt; 0 &amp;&amp; n &lt; HISTSZ) {
			d = h-&gt;line - a[n].incl-&gt;line;
			a[n-1].ldel += d;
			a[n-1].idel += d;
		}
	}

	if(n &gt; HISTSZ)
		n = HISTSZ;

	str[0] = 0;
	for(i=n-1; i&gt;=0; i--) {
		if(i != n-1) {
			if(fp-&gt;flags &amp; ~(FmtWidth|FmtPrec))
				break;
			strcat(str, &#34; &#34;);
		}
		if(a[i].line)
			snprint(s, STRINGSZ, &#34;%s:%ld[%s:%ld]&#34;,
				a[i].line-&gt;name, lno-a[i].ldel+1,
				a[i].incl-&gt;name, lno-a[i].idel+1);
		else
			snprint(s, STRINGSZ, &#34;%s:%ld&#34;,
				a[i].incl-&gt;name, lno-a[i].idel+1);
		if(strlen(s)+strlen(str) &gt;= STRINGSZ-10)
			break;
		strcat(str, s);
		lno = a[i].incl-&gt;line - 1;	/* now print out start of this file */
	}
	if(n == 0)
		strcat(str, &#34;&lt;epoch&gt;&#34;);

	return fmtstrcpy(fp, str);
}

/*
s%,%,\n%g
s%\n+%\n%g
s%^[ 	]*T%%g
s%,.*%%g
s%.+%	[T&amp;]		= &#34;&amp;&#34;,%g
s%^	........*\]%&amp;~%g
s%~	%%g
*/

static char*
etnames[] =
{
	[TINT]		= &#34;INT&#34;,
	[TUINT]		= &#34;UINT&#34;,
	[TINT8]		= &#34;INT8&#34;,
	[TUINT8]	= &#34;UINT8&#34;,
	[TINT16]	= &#34;INT16&#34;,
	[TUINT16]	= &#34;UINT16&#34;,
	[TINT32]	= &#34;INT32&#34;,
	[TUINT32]	= &#34;UINT32&#34;,
	[TINT64]	= &#34;INT64&#34;,
	[TUINT64]	= &#34;UINT64&#34;,
	[TUINTPTR]	= &#34;UINTPTR&#34;,
	[TFLOAT]	= &#34;FLOAT&#34;,
	[TFLOAT32]	= &#34;FLOAT32&#34;,
	[TFLOAT64]	= &#34;FLOAT64&#34;,
	[TBOOL]		= &#34;BOOL&#34;,
	[TPTR32]	= &#34;PTR32&#34;,
	[TPTR64]	= &#34;PTR64&#34;,
	[TDDD]		= &#34;DDD&#34;,
	[TFUNC]		= &#34;FUNC&#34;,
	[TARRAY]	= &#34;ARRAY&#34;,
	[TSTRUCT]	= &#34;STRUCT&#34;,
	[TCHAN]		= &#34;CHAN&#34;,
	[TMAP]		= &#34;MAP&#34;,
	[TINTER]	= &#34;INTER&#34;,
	[TFORW]		= &#34;FORW&#34;,
	[TFIELD]	= &#34;FIELD&#34;,
	[TSTRING]	= &#34;STRING&#34;,
	[TCHAN]		= &#34;CHAN&#34;,
	[TANY]		= &#34;ANY&#34;,
};

int
Econv(Fmt *fp)
{
	char buf[500];
	int et;

	et = va_arg(fp-&gt;args, int);
	if(et &lt; 0 || et &gt;= nelem(etnames) || etnames[et] == nil) {
		snprint(buf, sizeof(buf), &#34;E-%d&#34;, et);
		return fmtstrcpy(fp, buf);
	}
	return fmtstrcpy(fp, etnames[et]);
}

int
Jconv(Fmt *fp)
{
	Node *n;

	n = va_arg(fp-&gt;args, Node*);
	if(n-&gt;ullman != 0)
		fmtprint(fp, &#34; u(%d)&#34;, n-&gt;ullman);

	if(n-&gt;addable != 0)
		fmtprint(fp, &#34; a(%d)&#34;, n-&gt;addable);

	if(n-&gt;vargen != 0)
		fmtprint(fp, &#34; g(%ld)&#34;, n-&gt;vargen);

	if(n-&gt;lineno != 0)
		fmtprint(fp, &#34; l(%ld)&#34;, n-&gt;lineno);

	if(n-&gt;xoffset != 0)
		fmtprint(fp, &#34; x(%lld)&#34;, n-&gt;xoffset);

	if(n-&gt;class != 0)
		fmtprint(fp, &#34; class(%d)&#34;, n-&gt;class);

	if(n-&gt;colas != 0)
		fmtprint(fp, &#34; colas(%d)&#34;, n-&gt;colas);

	if(n-&gt;funcdepth != 0)
		fmtprint(fp, &#34; f(%d)&#34;, n-&gt;funcdepth);

	if(n-&gt;typecheck != 0)
		fmtprint(fp, &#34; tc(%d)&#34;, n-&gt;typecheck);

	if(n-&gt;dodata != 0)
		fmtprint(fp, &#34; dd(%d)&#34;, n-&gt;dodata);

	return 0;
}

int
Gconv(Fmt *fp)
{
	char buf[100];
	Type *t;

	t = va_arg(fp-&gt;args, Type*);

	if(t-&gt;etype == TFUNC) {
		if(t-&gt;vargen != 0) {
			snprint(buf, sizeof(buf), &#34;-%d%d%d g(%ld)&#34;,
				t-&gt;thistuple, t-&gt;outtuple, t-&gt;intuple, t-&gt;vargen);
			goto out;
		}
		snprint(buf, sizeof(buf), &#34;-%d%d%d&#34;,
			t-&gt;thistuple, t-&gt;outtuple, t-&gt;intuple);
		goto out;
	}
	if(t-&gt;vargen != 0) {
		snprint(buf, sizeof(buf), &#34; g(%ld)&#34;, t-&gt;vargen);
		goto out;
	}
	strcpy(buf, &#34;&#34;);

out:
	return fmtstrcpy(fp, buf);
}

int
Sconv(Fmt *fp)
{
	Sym *s;
	char *pkg, *nam;

	s = va_arg(fp-&gt;args, Sym*);
	if(s == S) {
		fmtstrcpy(fp, &#34;&lt;S&gt;&#34;);
		return 0;
	}

	pkg = &#34;&lt;nil&gt;&#34;;
	nam = pkg;

	if(s-&gt;package != nil)
		pkg = s-&gt;package;
	if(s-&gt;name != nil)
		nam = s-&gt;name;

	if(!(fp-&gt;flags &amp; FmtShort))
	if(strcmp(pkg, package) != 0 || (fp-&gt;flags &amp; FmtLong)) {
		fmtprint(fp, &#34;%s.%s&#34;, pkg, nam);
		return 0;
	}
	fmtstrcpy(fp, nam);
	return 0;
}

static char*
basicnames[] =
{
	[TINT]		= &#34;int&#34;,
	[TUINT]		= &#34;uint&#34;,
	[TINT8]		= &#34;int8&#34;,
	[TUINT8]	= &#34;uint8&#34;,
	[TINT16]	= &#34;int16&#34;,
	[TUINT16]	= &#34;uint16&#34;,
	[TINT32]	= &#34;int32&#34;,
	[TUINT32]	= &#34;uint32&#34;,
	[TINT64]	= &#34;int64&#34;,
	[TUINT64]	= &#34;uint64&#34;,
	[TUINTPTR]	= &#34;uintptr&#34;,
	[TFLOAT]	= &#34;float&#34;,
	[TFLOAT32]	= &#34;float32&#34;,
	[TFLOAT64]	= &#34;float64&#34;,
	[TBOOL]		= &#34;bool&#34;,
	[TANY]		= &#34;any&#34;,
	[TDDD]		= &#34;...&#34;,
	[TSTRING]		= &#34;string&#34;,
	[TNIL]		= &#34;nil&#34;,
	[TIDEAL]		= &#34;ideal&#34;,
	[TBLANK]		= &#34;blank&#34;,
};

int
Tpretty(Fmt *fp, Type *t)
{
	Type *t1;
	Sym *s;

	if(t-&gt;etype != TFIELD
	&amp;&amp; t-&gt;sym != S
	&amp;&amp; !(fp-&gt;flags&amp;FmtLong)) {
		s = t-&gt;sym;
		if(t == types[t-&gt;etype])
			return fmtprint(fp, &#34;%s&#34;, s-&gt;name);
		if(exporting) {
			if(fp-&gt;flags &amp; FmtShort)
				fmtprint(fp, &#34;%hS&#34;, s);
			else
				fmtprint(fp, &#34;%lS&#34;, s);
			if(strcmp(s-&gt;package, package) != 0)
				return 0;
			if(s-&gt;flags &amp; SymImported)
				return 0;
			if(t-&gt;vargen)
				fmtprint(fp, &#34;·%d&#34;, t-&gt;vargen);
			return 0;
		}
		return fmtprint(fp, &#34;%S&#34;, s);
	}

	if(t-&gt;etype &lt; nelem(basicnames) &amp;&amp; basicnames[t-&gt;etype] != nil) {
		if(isideal(t) &amp;&amp; t-&gt;etype != TIDEAL &amp;&amp; t-&gt;etype != TNIL)
			fmtprint(fp, &#34;ideal &#34;);
		return fmtprint(fp, &#34;%s&#34;, basicnames[t-&gt;etype]);
	}

	switch(t-&gt;etype) {
	case TPTR32:
	case TPTR64:
		if(fp-&gt;flags&amp;FmtShort)	// pass flag thru for methodsym
			return fmtprint(fp, &#34;*%hT&#34;, t-&gt;type);
		return fmtprint(fp, &#34;*%T&#34;, t-&gt;type);

	case TCHAN:
		switch(t-&gt;chan) {
		case Crecv:
			return fmtprint(fp, &#34;&lt;-chan %T&#34;, t-&gt;type);
		case Csend:
			if(t-&gt;type != T &amp;&amp; t-&gt;type-&gt;etype == TCHAN)
				return fmtprint(fp, &#34;chan&lt;- (%T)&#34;, t-&gt;type);
			return fmtprint(fp, &#34;chan&lt;- %T&#34;, t-&gt;type);
		}
		return fmtprint(fp, &#34;chan %T&#34;, t-&gt;type);

	case TMAP:
		return fmtprint(fp, &#34;map[%T] %T&#34;, t-&gt;down, t-&gt;type);

	case TFUNC:
		// t-&gt;type is method struct
		// t-&gt;type-&gt;down is result struct
		// t-&gt;type-&gt;down-&gt;down is arg struct
		if(t-&gt;thistuple &amp;&amp; !(fp-&gt;flags&amp;FmtSharp) &amp;&amp; !(fp-&gt;flags&amp;FmtShort)) {
			fmtprint(fp, &#34;method(&#34;);
			for(t1=getthisx(t)-&gt;type; t1; t1=t1-&gt;down) {
				fmtprint(fp, &#34;%T&#34;, t1);
				if(t1-&gt;down)
					fmtprint(fp, &#34;, &#34;);
			}
			fmtprint(fp, &#34;)&#34;);
		}

		if(!(fp-&gt;flags&amp;FmtByte))
			fmtprint(fp, &#34;func&#34;);
		fmtprint(fp, &#34;(&#34;);
		for(t1=getinargx(t)-&gt;type; t1; t1=t1-&gt;down) {
			if(noargnames &amp;&amp; t1-&gt;etype == TFIELD)
				fmtprint(fp, &#34;%T&#34;, t1-&gt;type);
			else
				fmtprint(fp, &#34;%T&#34;, t1);
			if(t1-&gt;down)
				fmtprint(fp, &#34;, &#34;);
		}
		fmtprint(fp, &#34;)&#34;);
		switch(t-&gt;outtuple) {
		case 0:
			break;
		case 1:
			t1 = getoutargx(t)-&gt;type;
			if(t1 == T) {
				// failure to typecheck earlier; don&#39;t know the type
				fmtprint(fp, &#34; ?unknown-type?&#34;);
				break;
			}
			if(t1-&gt;etype != TFIELD &amp;&amp; t1-&gt;etype != TFUNC) {
				fmtprint(fp, &#34; %T&#34;, t1);
				break;
			}
		default:
			t1 = getoutargx(t)-&gt;type;
			fmtprint(fp, &#34; (&#34;);
			for(; t1; t1=t1-&gt;down) {
				if(noargnames &amp;&amp; t1-&gt;etype == TFIELD)
					fmtprint(fp, &#34;%T&#34;, t1-&gt;type);
				else
					fmtprint(fp, &#34;%T&#34;, t1);
				if(t1-&gt;down)
					fmtprint(fp, &#34;, &#34;);
			}
			fmtprint(fp, &#34;)&#34;);
			break;
		}
		return 0;

	case TARRAY:
		if(t-&gt;bound &gt;= 0)
			return fmtprint(fp, &#34;[%d]%T&#34;, (int)t-&gt;bound, t-&gt;type);
		if(t-&gt;bound == -100)
			return fmtprint(fp, &#34;[...]%T&#34;, t-&gt;type);
		return fmtprint(fp, &#34;[]%T&#34;, t-&gt;type);

	case TINTER:
		fmtprint(fp, &#34;interface {&#34;);
		for(t1=t-&gt;type; t1!=T; t1=t1-&gt;down) {
			fmtprint(fp, &#34; %hS %hhT&#34;, t1-&gt;sym, t1-&gt;type);
			if(t1-&gt;down)
				fmtprint(fp, &#34;;&#34;);
		}
		return fmtprint(fp, &#34; }&#34;);

	case TSTRUCT:
		if(t-&gt;funarg) {
			fmtprint(fp, &#34;(&#34;);
			for(t1=t-&gt;type; t1!=T; t1=t1-&gt;down) {
				fmtprint(fp, &#34;%T&#34;, t1);
				if(t1-&gt;down)
					fmtprint(fp, &#34;, &#34;);
			}
			return fmtprint(fp, &#34;)&#34;);
		}
		fmtprint(fp, &#34;struct {&#34;);
		for(t1=t-&gt;type; t1!=T; t1=t1-&gt;down) {
			fmtprint(fp, &#34; %T&#34;, t1);
			if(t1-&gt;down)
				fmtprint(fp, &#34;;&#34;);
		}
		return fmtprint(fp, &#34; }&#34;);

	case TFIELD:
		if(t-&gt;sym == S || t-&gt;embedded) {
			if(exporting)
				fmtprint(fp, &#34;? &#34;);
			fmtprint(fp, &#34;%T&#34;, t-&gt;type);
		} else
			fmtprint(fp, &#34;%hS %T&#34;, t-&gt;sym, t-&gt;type);
		if(t-&gt;note)
			fmtprint(fp, &#34; \&#34;%Z\&#34;&#34;, t-&gt;note);
		return 0;

	case TFORW:
		if(exporting)
			yyerror(&#34;undefined type %S&#34;, t-&gt;sym);
		if(t-&gt;sym)
			return fmtprint(fp, &#34;undefined %S&#34;, t-&gt;sym);
		return fmtprint(fp, &#34;undefined&#34;);
	}

	// Don&#39;t know how to handle - fall back to detailed prints.
	return -1;
}

int
Tconv(Fmt *fp)
{
	Type *t, *t1;
	int r, et, sharp, minus;

	sharp = (fp-&gt;flags &amp; FmtSharp);
	minus = (fp-&gt;flags &amp; FmtLeft);
	fp-&gt;flags &amp;= ~(FmtSharp|FmtLeft);

	t = va_arg(fp-&gt;args, Type*);
	if(t == T)
		return fmtstrcpy(fp, &#34;&lt;T&gt;&#34;);

	t-&gt;trecur++;
	if(t-&gt;trecur &gt; 5) {
		fmtprint(fp, &#34;...&#34;);
		goto out;
	}

	if(!debug[&#39;t&#39;]) {
		if(sharp)
			exporting++;
		if(minus)
			noargnames++;
		r = Tpretty(fp, t);
		if(sharp)
			exporting--;
		if(minus)
			noargnames--;
		if(r &gt;= 0) {
			t-&gt;trecur--;
			return 0;
		}
	}

	et = t-&gt;etype;
	fmtprint(fp, &#34;%E &#34;, et);
	if(t-&gt;sym != S)
		fmtprint(fp, &#34;&lt;%S&gt;&#34;, t-&gt;sym);

	switch(et) {
	default:
		if(t-&gt;type != T)
			fmtprint(fp, &#34; %T&#34;, t-&gt;type);
		break;

	case TFIELD:
		fmtprint(fp, &#34;%T&#34;, t-&gt;type);
		break;

	case TFUNC:
		if(fp-&gt;flags &amp; FmtLong)
			fmtprint(fp, &#34;%d%d%d(%lT,%lT)%lT&#34;,
				t-&gt;thistuple, t-&gt;intuple, t-&gt;outtuple,
				t-&gt;type, t-&gt;type-&gt;down-&gt;down, t-&gt;type-&gt;down);
		else
			fmtprint(fp, &#34;%d%d%d(%T,%T)%T&#34;,
				t-&gt;thistuple, t-&gt;intuple, t-&gt;outtuple,
				t-&gt;type, t-&gt;type-&gt;down-&gt;down, t-&gt;type-&gt;down);
		break;

	case TINTER:
		fmtprint(fp, &#34;{&#34;);
		if(fp-&gt;flags &amp; FmtLong)
			for(t1=t-&gt;type; t1!=T; t1=t1-&gt;down)
				fmtprint(fp, &#34;%lT;&#34;, t1);
		fmtprint(fp, &#34;}&#34;);
		break;

	case TSTRUCT:
		fmtprint(fp, &#34;{&#34;);
		if(fp-&gt;flags &amp; FmtLong)
			for(t1=t-&gt;type; t1!=T; t1=t1-&gt;down)
				fmtprint(fp, &#34;%lT;&#34;, t1);
		fmtprint(fp, &#34;}&#34;);
		break;

	case TMAP:
		fmtprint(fp, &#34;[%T]%T&#34;, t-&gt;down, t-&gt;type);
		break;

	case TARRAY:
		if(t-&gt;bound &gt;= 0)
			fmtprint(fp, &#34;[%ld]%T&#34;, t-&gt;bound, t-&gt;type);
		else
			fmtprint(fp, &#34;[]%T&#34;, t-&gt;type);
		break;

	case TPTR32:
	case TPTR64:
		fmtprint(fp, &#34;%T&#34;, t-&gt;type);
		break;
	}

out:
	t-&gt;trecur--;
	return 0;
}

int
Nconv(Fmt *fp)
{
	char buf1[500];
	Node *n;

	n = va_arg(fp-&gt;args, Node*);
	if(n == N) {
		fmtprint(fp, &#34;&lt;N&gt;&#34;);
		goto out;
	}

	if(fp-&gt;flags &amp; FmtSign) {
		if(n-&gt;type == T)
			fmtprint(fp, &#34;%#N&#34;, n);
		else if(n-&gt;type-&gt;etype == TNIL)
			fmtprint(fp, &#34;nil&#34;);
		else
			fmtprint(fp, &#34;%#N (type %T)&#34;, n, n-&gt;type);
		goto out;
	}

	if(fp-&gt;flags &amp; FmtSharp) {
		exprfmt(fp, n, 0);
		goto out;
	}

	switch(n-&gt;op) {
	default:
		fmtprint(fp, &#34;%O%J&#34;, n-&gt;op, n);
		break;

	case ONAME:
	case ONONAME:
		if(n-&gt;sym == S) {
			fmtprint(fp, &#34;%O%J&#34;, n-&gt;op, n);
			break;
		}
		fmtprint(fp, &#34;%O-%S G%ld%J&#34;, n-&gt;op,
			n-&gt;sym, n-&gt;vargen, n);
		goto ptyp;

	case OREGISTER:
		fmtprint(fp, &#34;%O-%R%J&#34;, n-&gt;op, n-&gt;val.u.reg, n);
		break;

	case OLITERAL:
		switch(n-&gt;val.ctype) {
		default:
			snprint(buf1, sizeof(buf1), &#34;LITERAL-ctype=%d&#34;, n-&gt;val.ctype);
			break;
		case CTINT:
			snprint(buf1, sizeof(buf1), &#34;I%B&#34;, n-&gt;val.u.xval);
			break;
		case CTFLT:
			snprint(buf1, sizeof(buf1), &#34;F%g&#34;, mpgetflt(n-&gt;val.u.fval));
			break;
		case CTSTR:
			snprint(buf1, sizeof(buf1), &#34;S\&#34;%Z\&#34;&#34;, n-&gt;val.u.sval);
			break;
		case CTBOOL:
			snprint(buf1, sizeof(buf1), &#34;B%d&#34;, n-&gt;val.u.bval);
			break;
		case CTNIL:
			snprint(buf1, sizeof(buf1), &#34;N&#34;);
			break;
		}
		fmtprint(fp, &#34;%O-%s%J&#34;, n-&gt;op, buf1, n);
		break;

	case OASOP:
		fmtprint(fp, &#34;%O-%O%J&#34;, n-&gt;op, n-&gt;etype, n);
		break;

	case OTYPE:
		fmtprint(fp, &#34;%O %T&#34;, n-&gt;op, n-&gt;type);
		break;
	}
	if(n-&gt;sym != S)
		fmtprint(fp, &#34; %S G%ld&#34;, n-&gt;sym, n-&gt;vargen);

ptyp:
	if(n-&gt;type != T)
		fmtprint(fp, &#34; %T&#34;, n-&gt;type);

out:
	return 0;
}

Node*
treecopy(Node *n)
{
	Node *m;

	if(n == N)
		return N;

	switch(n-&gt;op) {
	default:
		m = nod(OXXX, N, N);
		*m = *n;
		m-&gt;left = treecopy(n-&gt;left);
		m-&gt;right = treecopy(n-&gt;right);
		m-&gt;list = listtreecopy(n-&gt;list);
		if(m-&gt;defn)
			abort();
		break;

	case ONONAME:
		if(n-&gt;iota) {
			m = nod(OIOTA, n, nodintconst(iota));
			break;
		}
		// fall through
	case OLITERAL:
	case ONAME:
	case OTYPE:
		m = n;
		break;
	}
	return m;
}

int
Zconv(Fmt *fp)
{
	Rune r;
	Strlit *sp;
	char *s, *se;

	sp = va_arg(fp-&gt;args, Strlit*);
	if(sp == nil)
		return fmtstrcpy(fp, &#34;&lt;nil&gt;&#34;);

	s = sp-&gt;s;
	se = s + sp-&gt;len;
	while(s &lt; se) {
		s += chartorune(&amp;r, s);
		switch(r) {
		default:
			if(r &lt; &#39; &#39;) {
				fmtprint(fp, &#34;\\x%02x&#34;, r);
				break;
			}
			fmtrune(fp, r);
			break;
		case &#39;\t&#39;:
			fmtstrcpy(fp, &#34;\\t&#34;);
			break;
		case &#39;\n&#39;:
			fmtstrcpy(fp, &#34;\\n&#34;);
			break;
		case &#39;\&#34;&#39;:
		case &#39;\\&#39;:
			fmtrune(fp, &#39;\\&#39;);
			fmtrune(fp, r);
			break;
		}
	}
	return 0;
}

int
isnil(Node *n)
{
	if(n == N)
		return 0;
	if(n-&gt;op != OLITERAL)
		return 0;
	if(n-&gt;val.ctype != CTNIL)
		return 0;
	return 1;
}

int
isptrto(Type *t, int et)
{
	if(t == T)
		return 0;
	if(!isptr[t-&gt;etype])
		return 0;
	t = t-&gt;type;
	if(t == T)
		return 0;
	if(t-&gt;etype != et)
		return 0;
	return 1;
}

int
istype(Type *t, int et)
{
	return t != T &amp;&amp; t-&gt;etype == et;
}

int
isfixedarray(Type *t)
{
	return t != T &amp;&amp; t-&gt;etype == TARRAY &amp;&amp; t-&gt;bound &gt;= 0;
}

int
isslice(Type *t)
{
	return t != T &amp;&amp; t-&gt;etype == TARRAY &amp;&amp; t-&gt;bound &lt; 0;
}

int
isblank(Node *n)
{
	char *p;

	if(n == N || n-&gt;sym == S)
		return 0;
	p = n-&gt;sym-&gt;name;
	if(p == nil)
		return 0;
	return p[0] == &#39;_&#39; &amp;&amp; p[1] == &#39;\0&#39;;
}

int
isselect(Node *n)
{
	Sym *s;

	if(n == N)
		return 0;
	n = n-&gt;left;
	s = pkglookup(&#34;selectsend&#34;, &#34;runtime&#34;);
	if(s == n-&gt;sym)
		return 1;
	s = pkglookup(&#34;selectrecv&#34;, &#34;runtime&#34;);
	if(s == n-&gt;sym)
		return 1;
	s = pkglookup(&#34;selectdefault&#34;, &#34;runtime&#34;);
	if(s == n-&gt;sym)
		return 1;
	return 0;
}

int
isinter(Type *t)
{
	if(t != T) {
		if(t-&gt;etype == TINTER)
			return 1;
		if(t-&gt;etype == TDDD)
			return 1;
	}
	return 0;
}

int
isnilinter(Type *t)
{
	if(!isinter(t))
		return 0;
	if(t-&gt;type != T)
		return 0;
	return 1;
}

int
isddd(Type *t)
{
	if(t != T &amp;&amp; t-&gt;etype == TDDD)
		return 1;
	return 0;
}

int
isideal(Type *t)
{
	if(t == T)
		return 0;
	return t == idealstring || t == idealbool || t-&gt;etype == TNIL || t-&gt;etype == TIDEAL;
}

/*
 * given receiver of type t (t == r or t == *r)
 * return type to hang methods off (r).
 */
Type*
methtype(Type *t)
{
	int ptr;

	if(t == T)
		return T;

	// strip away pointer if it&#39;s there
	ptr = 0;
	if(isptr[t-&gt;etype]) {
		if(t-&gt;sym != S)
			return T;
		ptr = 1;
		t = t-&gt;type;
		if(t == T)
			return T;
	}

	// need a type name
	if(t-&gt;sym == S)
		return T;

	// check types
	if(!issimple[t-&gt;etype])
	switch(t-&gt;etype) {
	default:
		return T;
	case TSTRUCT:
	case TARRAY:
	case TMAP:
	case TCHAN:
	case TSTRING:
	case TFUNC:
		break;
	}

	return t;
}

int
iscomposite(Type *t)
{
	if(t == T)
		return 0;
	switch(t-&gt;etype) {
	case TARRAY:
	case TSTRUCT:
	case TMAP:
		return 1;
	}
	return 0;
}

int
eqtype1(Type *t1, Type *t2, int d, int names)
{
	if(d &gt;= 20)
		return 1;
	if(t1 == t2)
		return 1;
	if(t1 == T || t2 == T)
		return 0;
	if(t1-&gt;etype != t2-&gt;etype)
		return 0;
	if(names &amp;&amp; t1-&gt;etype != TFIELD &amp;&amp; t1-&gt;sym &amp;&amp; t2-&gt;sym &amp;&amp; t1 != t2)
		return 0;
	switch(t1-&gt;etype) {
	case TINTER:
	case TSTRUCT:
		t1 = t1-&gt;type;
		t2 = t2-&gt;type;
		for(;;) {
			if(!eqtype1(t1, t2, d+1, names))
				return 0;
			if(t1 == T)
				return 1;
			if(t1-&gt;embedded != t2-&gt;embedded)
				return 0;
			if(t1-&gt;nname != N &amp;&amp; t1-&gt;nname-&gt;sym != S) {
				if(t2-&gt;nname == N || t2-&gt;nname-&gt;sym == S)
					return 0;
				if(strcmp(t1-&gt;nname-&gt;sym-&gt;name, t2-&gt;nname-&gt;sym-&gt;name) != 0)
					return 0;
			}
			t1 = t1-&gt;down;
			t2 = t2-&gt;down;
		}
		return 1;

	case TFUNC:
		// Loop over structs: receiver, in, out.
		t1 = t1-&gt;type;
		t2 = t2-&gt;type;
		for(;;) {
			Type *ta, *tb;
			if(t1 == t2)
				break;
			if(t1 == T || t2 == T)
				return 0;
			if(t1-&gt;etype != TSTRUCT || t2-&gt;etype != TSTRUCT)
				return 0;

			// Loop over fields in structs, checking type only.
			ta = t1-&gt;type;
			tb = t2-&gt;type;
			while(ta != tb) {
				if(ta == T || tb == T)
					return 0;
				if(ta-&gt;etype != TFIELD || tb-&gt;etype != TFIELD)
					return 0;
				if(!eqtype1(ta-&gt;type, tb-&gt;type, d+1, names))
					return 0;
				ta = ta-&gt;down;
				tb = tb-&gt;down;
			}

			t1 = t1-&gt;down;
			t2 = t2-&gt;down;
		}
		return 1;

	case TARRAY:
		if(t1-&gt;bound != t2-&gt;bound)
			return 0;
		break;

	case TCHAN:
		if(t1-&gt;chan != t2-&gt;chan)
			return 0;
		break;

	case TMAP:
		if(!eqtype1(t1-&gt;down, t2-&gt;down, d+1, names))
			return 0;
		break;
	}
	return eqtype1(t1-&gt;type, t2-&gt;type, d+1, names);
}

int
eqtype(Type *t1, Type *t2)
{
	return eqtype1(t1, t2, 0, 1);
}

/*
 * can we convert from type src to dst with
 * a trivial conversion (no bits changing)?
 */
int
cvttype(Type *dst, Type *src)
{
	return eqtype1(dst, src, 0, 0);
}

int
eqtypenoname(Type *t1, Type *t2)
{
	if(t1 == T || t2 == T || t1-&gt;etype != TSTRUCT || t2-&gt;etype != TSTRUCT)
		return eqtype(t1, t2);

	t1 = t1-&gt;type;
	t2 = t2-&gt;type;
	for(;;) {
		if(!eqtype(t1, t2))
			return 0;
		if(t1 == T)
			return 1;
		t1 = t1-&gt;down;
		t2 = t2-&gt;down;
	}
}

static int
subtype(Type **stp, Type *t, int d)
{
	Type *st;

loop:
	st = *stp;
	if(st == T)
		return 0;

	d++;
	if(d &gt;= 10)
		return 0;

	switch(st-&gt;etype) {
	default:
		return 0;

	case TPTR32:
	case TPTR64:
	case TCHAN:
	case TARRAY:
		stp = &amp;st-&gt;type;
		goto loop;

	case TANY:
		if(!st-&gt;copyany)
			return 0;
		*stp = t;
		break;

	case TMAP:
		if(subtype(&amp;st-&gt;down, t, d))
			break;
		stp = &amp;st-&gt;type;
		goto loop;

	case TFUNC:
		for(;;) {
			if(subtype(&amp;st-&gt;type, t, d))
				break;
			if(subtype(&amp;st-&gt;type-&gt;down-&gt;down, t, d))
				break;
			if(subtype(&amp;st-&gt;type-&gt;down, t, d))
				break;
			return 0;
		}
		break;

	case TSTRUCT:
		for(st=st-&gt;type; st!=T; st=st-&gt;down)
			if(subtype(&amp;st-&gt;type, t, d))
				return 1;
		return 0;
	}
	return 1;
}

/*
 * Is this a 64-bit type?
 */
int
is64(Type *t)
{
	if(t == T)
		return 0;
	switch(simtype[t-&gt;etype]) {
	case TINT64:
	case TUINT64:
	case TPTR64:
		return 1;
	}
	return 0;
}

/*
 * Is a conversion between t1 and t2 a no-op?
 */
int
noconv(Type *t1, Type *t2)
{
	int e1, e2;

	e1 = simtype[t1-&gt;etype];
	e2 = simtype[t2-&gt;etype];

	switch(e1) {
	case TINT8:
	case TUINT8:
		return e2 == TINT8 || e2 == TUINT8;

	case TINT16:
	case TUINT16:
		return e2 == TINT16 || e2 == TUINT16;

	case TINT32:
	case TUINT32:
	case TPTR32:
		return e2 == TINT32 || e2 == TUINT32 || e2 == TPTR32;

	case TINT64:
	case TUINT64:
	case TPTR64:
		return e2 == TINT64 || e2 == TUINT64 || e2 == TPTR64;

	case TFLOAT32:
		return e2 == TFLOAT32;

	case TFLOAT64:
		return e2 == TFLOAT64;
	}
	return 0;
}

void
argtype(Node *on, Type *t)
{
	dowidth(t);
	if(!subtype(&amp;on-&gt;type, t, 0))
		fatal(&#34;argtype: failed %N %T\n&#34;, on, t);
}

Type*
shallow(Type *t)
{
	Type *nt;

	if(t == T)
		return T;
	nt = typ(0);
	*nt = *t;
	return nt;
}

Type*
deep(Type *t)
{
	Type *nt, *xt;

	if(t == T)
		return T;

	switch(t-&gt;etype) {
	default:
		nt = t;	// share from here down
		break;

	case TANY:
		nt = shallow(t);
		nt-&gt;copyany = 1;
		break;

	case TPTR32:
	case TPTR64:
	case TCHAN:
	case TARRAY:
		nt = shallow(t);
		nt-&gt;type = deep(t-&gt;type);
		break;

	case TMAP:
		nt = shallow(t);
		nt-&gt;down = deep(t-&gt;down);
		nt-&gt;type = deep(t-&gt;type);
		break;

	case TFUNC:
		nt = shallow(t);
		nt-&gt;type = deep(t-&gt;type);
		nt-&gt;type-&gt;down = deep(t-&gt;type-&gt;down);
		nt-&gt;type-&gt;down-&gt;down = deep(t-&gt;type-&gt;down-&gt;down);
		break;

	case TSTRUCT:
		nt = shallow(t);
		nt-&gt;type = shallow(t-&gt;type);
		xt = nt-&gt;type;

		for(t=t-&gt;type; t!=T; t=t-&gt;down) {
			xt-&gt;type = deep(t-&gt;type);
			xt-&gt;down = shallow(t-&gt;down);
			xt = xt-&gt;down;
		}
		break;
	}
	return nt;
}

Node*
syslook(char *name, int copy)
{
	Sym *s;
	Node *n;

	s = pkglookup(name, &#34;runtime&#34;);
	if(s == S || s-&gt;def == N)
		fatal(&#34;looksys: cant find runtime.%s&#34;, name);

	if(!copy)
		return s-&gt;def;

	n = nod(0, N, N);
	*n = *s-&gt;def;
	n-&gt;type = deep(s-&gt;def-&gt;type);

	return n;
}

/*
 * are the arg names of two
 * functions the same. we know
 * that eqtype has been called
 * and has returned true.
 */
int
eqargs(Type *t1, Type *t2)
{
	if(t1 == t2)
		return 1;
	if(t1 == T || t2 == T)
		return 0;

	if(t1-&gt;etype != t2-&gt;etype)
		return 0;

	if(t1-&gt;etype != TFUNC)
		fatal(&#34;eqargs: oops %E&#34;, t1-&gt;etype);

	t1 = t1-&gt;type;
	t2 = t2-&gt;type;
	for(;;) {
		if(t1 == t2)
			break;
		if(!eqtype(t1, t2))
			return 0;
		t1 = t1-&gt;down;
		t2 = t2-&gt;down;
	}
	return 1;
}

/*
 * compute a hash value for type t.
 * if t is a method type, ignore the receiver
 * so that the hash can be used in interface checks.
 * %#-T (which calls Tpretty, above) already contains
 * all the necessary logic to generate a representation
 * of the type that completely describes it.
 * using smprint here avoids duplicating that code.
 * using md5 here is overkill, but i got tired of
 * accidental collisions making the runtime think
 * two types are equal when they really aren&#39;t.
 */
uint32
typehash(Type *t)
{
	char *p;
	MD5 d;

	if(t-&gt;thistuple) {
		// hide method receiver from Tpretty
		t-&gt;thistuple = 0;
		p = smprint(&#34;%#-T&#34;, t);
		t-&gt;thistuple = 1;
	}else
		p = smprint(&#34;%#-T&#34;, t);
	md5reset(&amp;d);
	md5write(&amp;d, (uchar*)p, strlen(p));
	free(p);
	return md5sum(&amp;d);
}

Type*
ptrto(Type *t)
{
	Type *t1;

	if(tptr == 0)
		fatal(&#34;ptrto: nil&#34;);
	t1 = typ(tptr);
	t1-&gt;type = t;
	t1-&gt;width = types[tptr]-&gt;width;
	return t1;
}

void
frame(int context)
{
	char *p;
	NodeList *l;
	Node *n;
	int flag;

	p = &#34;stack&#34;;
	l = nil;
	if(curfn)
		l = curfn-&gt;dcl;
	if(context) {
		p = &#34;external&#34;;
		l = externdcl;
	}

	flag = 1;
	for(; l; l=l-&gt;next) {
		n = l-&gt;n;
		switch(n-&gt;op) {
		case ONAME:
			if(flag)
				print(&#34;--- %s frame ---\n&#34;, p);
			print(&#34;%O %S G%ld T\n&#34;, n-&gt;op, n-&gt;sym, n-&gt;vargen, n-&gt;type);
			flag = 0;
			break;

		case OTYPE:
			if(flag)
				print(&#34;--- %s frame ---\n&#34;, p);
			print(&#34;%O %T\n&#34;, n-&gt;op, n-&gt;type);
			flag = 0;
			break;
		}
	}
}

/*
 * calculate sethi/ullman number
 * roughly how many registers needed to
 * compile a node. used to compile the
 * hardest side first to minimize registers.
 */
void
ullmancalc(Node *n)
{
	int ul, ur;

	if(n == N)
		return;

	switch(n-&gt;op) {
	case OREGISTER:
	case OLITERAL:
	case ONAME:
		ul = 1;
		if(n-&gt;class == PPARAMREF || (n-&gt;class &amp; PHEAP))
			ul++;
		goto out;
	case OCALL:
	case OCALLFUNC:
	case OCALLMETH:
	case OCALLINTER:
		ul = UINF;
		goto out;
	}
	ul = 1;
	if(n-&gt;left != N)
		ul = n-&gt;left-&gt;ullman;
	ur = 1;
	if(n-&gt;right != N)
		ur = n-&gt;right-&gt;ullman;
	if(ul == ur)
		ul += 1;
	if(ur &gt; ul)
		ul = ur;

out:
	n-&gt;ullman = ul;
}

void
badtype(int o, Type *tl, Type *tr)
{
	yyerror(&#34;illegal types for operand: %O&#34;, o);
	if(tl != T)
		print(&#34;	%T\n&#34;, tl);
	if(tr != T)
		print(&#34;	%T\n&#34;, tr);

	// common mistake: *struct and *interface.
	if(tl &amp;&amp; tr &amp;&amp; isptr[tl-&gt;etype] &amp;&amp; isptr[tr-&gt;etype]) {
		if(tl-&gt;type-&gt;etype == TSTRUCT &amp;&amp; tr-&gt;type-&gt;etype == TINTER)
			print(&#34;	(*struct vs *interface)\n&#34;);
		else if(tl-&gt;type-&gt;etype == TINTER &amp;&amp; tr-&gt;type-&gt;etype == TSTRUCT)
			print(&#34;	(*interface vs *struct)\n&#34;);
	}
}

/*
 * iterator to walk a structure declaration
 */
Type*
structfirst(Iter *s, Type **nn)
{
	Type *n, *t;

	n = *nn;
	if(n == T)
		goto bad;

	switch(n-&gt;etype) {
	default:
		goto bad;

	case TSTRUCT:
	case TINTER:
	case TFUNC:
		break;
	}

	t = n-&gt;type;
	if(t == T)
		goto rnil;

	if(t-&gt;etype != TFIELD)
		fatal(&#34;structfirst: not field %T&#34;, t);

	s-&gt;t = t;
	return t;

bad:
	fatal(&#34;structfirst: not struct %T&#34;, n);

rnil:
	return T;
}

Type*
structnext(Iter *s)
{
	Type *n, *t;

	n = s-&gt;t;
	t = n-&gt;down;
	if(t == T)
		goto rnil;

	if(t-&gt;etype != TFIELD)
		goto bad;

	s-&gt;t = t;
	return t;

bad:
	fatal(&#34;structnext: not struct %T&#34;, n);

rnil:
	return T;
}

/*
 * iterator to this and inargs in a function
 */
Type*
funcfirst(Iter *s, Type *t)
{
	Type *fp;

	if(t == T)
		goto bad;

	if(t-&gt;etype != TFUNC)
		goto bad;

	s-&gt;tfunc = t;
	s-&gt;done = 0;
	fp = structfirst(s, getthis(t));
	if(fp == T) {
		s-&gt;done = 1;
		fp = structfirst(s, getinarg(t));
	}
	return fp;

bad:
	fatal(&#34;funcfirst: not func %T&#34;, t);
	return T;
}

Type*
funcnext(Iter *s)
{
	Type *fp;

	fp = structnext(s);
	if(fp == T &amp;&amp; !s-&gt;done) {
		s-&gt;done = 1;
		fp = structfirst(s, getinarg(s-&gt;tfunc));
	}
	return fp;
}

Type**
getthis(Type *t)
{
	if(t-&gt;etype != TFUNC)
		fatal(&#34;getthis: not a func %T&#34;, t);
	return &amp;t-&gt;type;
}

Type**
getoutarg(Type *t)
{
	if(t-&gt;etype != TFUNC)
		fatal(&#34;getoutarg: not a func %T&#34;, t);
	return &amp;t-&gt;type-&gt;down;
}

Type**
getinarg(Type *t)
{
	if(t-&gt;etype != TFUNC)
		fatal(&#34;getinarg: not a func %T&#34;, t);
	return &amp;t-&gt;type-&gt;down-&gt;down;
}

Type*
getthisx(Type *t)
{
	return *getthis(t);
}

Type*
getoutargx(Type *t)
{
	return *getoutarg(t);
}

Type*
getinargx(Type *t)
{
	return *getinarg(t);
}

/*
 * return !(op)
 * eg == &lt;=&gt; !=
 */
int
brcom(int a)
{
	switch(a) {
	case OEQ:	return ONE;
	case ONE:	return OEQ;
	case OLT:	return OGE;
	case OGT:	return OLE;
	case OLE:	return OGT;
	case OGE:	return OLT;
	}
	fatal(&#34;brcom: no com for %A\n&#34;, a);
	return a;
}

/*
 * return reverse(op)
 * eg a op b &lt;=&gt; b r(op) a
 */
int
brrev(int a)
{
	switch(a) {
	case OEQ:	return OEQ;
	case ONE:	return ONE;
	case OLT:	return OGT;
	case OGT:	return OLT;
	case OLE:	return OGE;
	case OGE:	return OLE;
	}
	fatal(&#34;brcom: no rev for %A\n&#34;, a);
	return a;
}

Node*
staticname(Type *t)
{
	Node *n;

	snprint(namebuf, sizeof(namebuf), &#34;statictmp_%.4d&#34;, statuniqgen);
	statuniqgen++;
	n = newname(lookup(namebuf));
	addvar(n, t, PEXTERN);
	return n;
}

/*
 * return side effect-free, assignable n, appending side effects to init.
 */
Node*
saferef(Node *n, NodeList **init)
{
	Node *l;
	Node *r;
	Node *a;

	switch(n-&gt;op) {
	case ONAME:
		return n;
	case ODOT:
		l = saferef(n-&gt;left, init);
		if(l == n-&gt;left)
			return n;
		r = nod(OXXX, N, N);
		*r = *n;
		r-&gt;left = l;
		typecheck(&amp;r, Erv);
		walkexpr(&amp;r, init);
		return r;

	case OINDEX:
	case ODOTPTR:
	case OIND:
		l = nod(OXXX, N, N);
		tempname(l, ptrto(n-&gt;type));
		a = nod(OAS, l, nod(OADDR, n, N));
		typecheck(&amp;a, Etop);
		walkexpr(&amp;a, init);
		*init = list(*init, a);
		r = nod(OIND, l, N);
		typecheck(&amp;r, Erv);
		walkexpr(&amp;r, init);
		return r;
	}
	fatal(&#34;saferef %N&#34;, n);
	return N;
}

/*
 * return side effect-free n, appending side effects to init.
 */
Node*
safeval(Node *n, NodeList **init)
{
	Node *l;
	Node *a;

	// is this a local variable or a dot of a local variable?
	for(l=n; l-&gt;op == ODOT; l=l-&gt;left)
		if(l-&gt;left-&gt;type != T &amp;&amp; isptr[l-&gt;left-&gt;type-&gt;etype])
			goto copy;
	if(l-&gt;op == ONAME &amp;&amp; (l-&gt;class == PAUTO || l-&gt;class == PPARAM))
		return n;

copy:
	l = nod(OXXX, N, N);
	tempname(l, n-&gt;type);
	a = nod(OAS, l, n);
	typecheck(&amp;a, Etop);
	walkexpr(&amp;a, init);
	*init = list(*init, a);
	return l;
}

void
setmaxarg(Type *t)
{
	int32 w;

	dowidth(t);
	w = t-&gt;argwid;
	if(t-&gt;argwid &gt;= 100000000)
		fatal(&#34;bad argwid %T&#34;, t);
	if(w &gt; maxarg)
		maxarg = w;
}

/*
 * code to resolve elided DOTs
 * in embedded types
 */

// search depth 0 --
// return count of fields+methods
// found with a given name
int
lookdot0(Sym *s, Type *t, Type **save)
{
	Type *f, *u;
	int c;

	u = t;
	if(isptr[u-&gt;etype])
		u = u-&gt;type;

	c = 0;
	if(u-&gt;etype == TSTRUCT || u-&gt;etype == TINTER) {
		for(f=u-&gt;type; f!=T; f=f-&gt;down)
			if(f-&gt;sym == s) {
				if(save)
					*save = f;
				c++;
			}
	}
	u = methtype(t);
	if(u != T) {
		for(f=u-&gt;method; f!=T; f=f-&gt;down)
			if(f-&gt;sym == s &amp;&amp; f-&gt;embedded == 0) {
				if(save)
					*save = f;
				c++;
			}
	}
	return c;
}

// search depth d --
// return count of fields+methods
// found at search depth.
// answer is in dotlist array and
// count of number of ways is returned.
int
adddot1(Sym *s, Type *t, int d, Type **save)
{
	Type *f, *u;
	int c, a;

	if(t-&gt;trecur)
		return 0;
	t-&gt;trecur = 1;

	if(d == 0) {
		c = lookdot0(s, t, save);
		goto out;
	}

	c = 0;
	u = t;
	if(isptr[u-&gt;etype])
		u = u-&gt;type;
	if(u-&gt;etype != TSTRUCT &amp;&amp; u-&gt;etype != TINTER)
		goto out;

	d--;
	for(f=u-&gt;type; f!=T; f=f-&gt;down) {
		if(!f-&gt;embedded)
			continue;
		if(f-&gt;sym == S)
			continue;
		a = adddot1(s, f-&gt;type, d, save);
		if(a != 0 &amp;&amp; c == 0)
			dotlist[d].field = f;
		c += a;
	}

out:
	t-&gt;trecur = 0;
	return c;
}

// in T.field
// find missing fields that
// will give shortest unique addressing.
// modify the tree with missing type names.
Node*
adddot(Node *n)
{
	Type *t;
	Sym *s;
	int c, d;

	typecheck(&amp;n-&gt;left, Erv);
	t = n-&gt;left-&gt;type;
	if(t == T)
		goto ret;

	if(n-&gt;right-&gt;op != ONAME)
		goto ret;
	s = n-&gt;right-&gt;sym;
	if(s == S)
		goto ret;

	for(d=0; d&lt;nelem(dotlist); d++) {
		c = adddot1(s, t, d, nil);
		if(c &gt; 0)
			goto out;
	}
	goto ret;

out:
	if(c &gt; 1)
		yyerror(&#34;ambiguous DOT reference %T.%S&#34;, t, s);

	// rebuild elided dots
	for(c=d-1; c&gt;=0; c--)
		n-&gt;left = nod(ODOT, n-&gt;left, newname(dotlist[c].field-&gt;sym));
ret:
	return n;
}


/*
 * code to help generate trampoline
 * functions for methods on embedded
 * subtypes.
 * these are approx the same as
 * the corresponding adddot routines
 * except that they expect to be called
 * with unique tasks and they return
 * the actual methods.
 */

typedef	struct	Symlink	Symlink;
struct	Symlink
{
	Type*		field;
	uchar		good;
	uchar		followptr;
	Symlink*	link;
};
static	Symlink*	slist;

static void
expand0(Type *t, int followptr)
{
	Type *f, *u;
	Symlink *sl;

	u = t;
	if(isptr[u-&gt;etype]) {
		followptr = 1;
		u = u-&gt;type;
	}

	if(u-&gt;etype == TINTER) {
		for(f=u-&gt;type; f!=T; f=f-&gt;down) {
			if(!exportname(f-&gt;sym-&gt;name) &amp;&amp; strcmp(f-&gt;sym-&gt;package, package) != 0)
				continue;
			if(f-&gt;sym-&gt;flags &amp; SymUniq)
				continue;
			f-&gt;sym-&gt;flags |= SymUniq;
			sl = mal(sizeof(*sl));
			sl-&gt;field = f;
			sl-&gt;link = slist;
			sl-&gt;followptr = followptr;
			slist = sl;
		}
		return;
	}

	u = methtype(t);
	if(u != T) {
		for(f=u-&gt;method; f!=T; f=f-&gt;down) {
			if(!exportname(f-&gt;sym-&gt;name) &amp;&amp; strcmp(f-&gt;sym-&gt;package, package) != 0)
				continue;
			if(f-&gt;sym-&gt;flags &amp; SymUniq)
				continue;
			f-&gt;sym-&gt;flags |= SymUniq;
			sl = mal(sizeof(*sl));
			sl-&gt;field = f;
			sl-&gt;link = slist;
			sl-&gt;followptr = followptr;
			slist = sl;
		}
	}
}

static void
expand1(Type *t, int d, int followptr)
{
	Type *f, *u;

	if(t-&gt;trecur)
		return;
	if(d == 0)
		return;
	t-&gt;trecur = 1;

	if(d != nelem(dotlist)-1)
		expand0(t, followptr);

	u = t;
	if(isptr[u-&gt;etype]) {
		followptr = 1;
		u = u-&gt;type;
	}
	if(u-&gt;etype != TSTRUCT &amp;&amp; u-&gt;etype != TINTER)
		goto out;

	for(f=u-&gt;type; f!=T; f=f-&gt;down) {
		if(!f-&gt;embedded)
			continue;
		if(f-&gt;sym == S)
			continue;
		expand1(f-&gt;type, d-1, followptr);
	}

out:
	t-&gt;trecur = 0;
}

void
expandmeth(Sym *s, Type *t)
{
	Symlink *sl;
	Type *f;
	int c, d;

	if(s == S)
		return;
	if(t == T || t-&gt;xmethod != nil)
		return;

	// generate all reachable methods
	slist = nil;
	expand1(t, nelem(dotlist)-1, 0);

	// check each method to be uniquely reachable
	for(sl=slist; sl!=nil; sl=sl-&gt;link) {
		sl-&gt;field-&gt;sym-&gt;flags &amp;= ~SymUniq;
		for(d=0; d&lt;nelem(dotlist); d++) {
			c = adddot1(sl-&gt;field-&gt;sym, t, d, &amp;f);
			if(c == 0)
				continue;
			if(c == 1) {
				sl-&gt;good = 1;
				sl-&gt;field = f;
			}
			break;
		}
	}

	t-&gt;xmethod = t-&gt;method;
	for(sl=slist; sl!=nil; sl=sl-&gt;link) {
		if(sl-&gt;good) {
			// add it to the base type method list
			f = typ(TFIELD);
			*f = *sl-&gt;field;
			f-&gt;embedded = 1;	// needs a trampoline
			if(sl-&gt;followptr)
				f-&gt;embedded = 2;
			f-&gt;down = t-&gt;xmethod;
			t-&gt;xmethod = f;

		}
	}
}

/*
 * Given funarg struct list, return list of ODCLFIELD Node fn args.
 */
NodeList*
structargs(Type **tl, int mustname)
{
	Iter savet;
	Node *a, *n;
	NodeList *args;
	Type *t;
	char buf[100];
	int gen;

	args = nil;
	gen = 0;
	for(t = structfirst(&amp;savet, tl); t != T; t = structnext(&amp;savet)) {
		n = N;
		if(t-&gt;sym)
			n = newname(t-&gt;sym);
		else if(mustname) {
			// have to give it a name so we can refer to it in trampoline
			snprint(buf, sizeof buf, &#34;.anon%d&#34;, gen++);
			n = newname(lookup(buf));
		}
		a = nod(ODCLFIELD, n, typenod(t-&gt;type));
		args = list(args, a);
	}
	return args;
}

/*
 * Generate a wrapper function to convert from
 * a receiver of type T to a receiver of type U.
 * That is,
 *
 *	func (t T) M() {
 *		...
 *	}
 *
 * already exists; this function generates
 *
 *	func (u U) M() {
 *		u.M()
 *	}
 *
 * where the types T and U are such that u.M() is valid
 * and calls the T.M method.
 * The resulting function is for use in method tables.
 *
 *	rcvr - U
 *	method - M func (t T)(), a TFIELD type struct
 *	newnam - the eventual mangled name of this function
 */
void
genwrapper(Type *rcvr, Type *method, Sym *newnam)
{
	Node *this, *fn, *call, *n, *t;
	NodeList *l, *args, *in, *out;

	if(debug[&#39;r&#39;])
		print(&#34;genwrapper rcvrtype=%T method=%T newnam=%S\n&#34;,
			rcvr, method, newnam);

	dclcontext = PEXTERN;
	markdcl();

	this = nod(ODCLFIELD, newname(lookup(&#34;.this&#34;)), typenod(rcvr));
	this-&gt;left-&gt;ntype = this-&gt;right;
	in = structargs(getinarg(method-&gt;type), 1);
	out = structargs(getoutarg(method-&gt;type), 0);

	fn = nod(ODCLFUNC, N, N);
	fn-&gt;nname = newname(newnam);
	t = nod(OTFUNC, this, N);
	t-&gt;list = in;
	t-&gt;rlist = out;
	fn-&gt;nname-&gt;ntype = t;
	funchdr(fn);

	// arg list
	args = nil;
	for(l=in; l; l=l-&gt;next)
		args = list(args, l-&gt;n-&gt;left);

	// generate call
	call = nod(OCALL, adddot(nod(OXDOT, this-&gt;left, newname(method-&gt;sym))), N);
	call-&gt;list = args;
	fn-&gt;nbody = list1(call);
	if(method-&gt;type-&gt;outtuple &gt; 0) {
		n = nod(ORETURN, N, N);
		n-&gt;list = fn-&gt;nbody;
		fn-&gt;nbody = list1(n);
	}

	if(debug[&#39;r&#39;])
		dumplist(&#34;genwrapper body&#34;, fn-&gt;nbody);

	funcbody(fn);
	typecheck(&amp;fn, Etop);
	funccompile(fn);
}

/*
 * delayed interface type check.
 * remember that there is an interface conversion
 * on the given line.  once the file is completely read
 * and all methods are known, we can check that
 * the conversions are valid.
 */

typedef struct Icheck Icheck;
struct Icheck
{
	Icheck *next;
	Type *dst;
	Type *src;
	int lineno;
	int explicit;
};
Icheck *icheck;
Icheck *ichecktail;

void
ifacecheck(Type *dst, Type *src, int lineno, int explicit)
{
	Icheck *p;

	p = mal(sizeof *p);
	if(ichecktail)
		ichecktail-&gt;next = p;
	else
		icheck = p;
	p-&gt;dst = dst;
	p-&gt;src = src;
	p-&gt;lineno = lineno;
	p-&gt;explicit = explicit;
	ichecktail = p;
}

Type*
ifacelookdot(Sym *s, Type *t, int *followptr)
{
	int i, c, d;
	Type *m;

	*followptr = 0;

	if(t == T)
		return T;

	for(d=0; d&lt;nelem(dotlist); d++) {
		c = adddot1(s, t, d, &amp;m);
		if(c &gt; 1) {
			yyerror(&#34;%T.%S is ambiguous&#34;, t, s);
			return T;
		}
		if(c == 1) {
			for(i=0; i&lt;d; i++) {
				if(isptr[dotlist[i].field-&gt;type-&gt;etype]) {
					*followptr = 1;
					break;
				}
			}
			return m;
		}
	}
	return T;
}

// check whether non-interface type t
// satisifes inteface type iface.
int
ifaceokT2I(Type *t0, Type *iface, Type **m, Type **samename)
{
	Type *t, *im, *tm, *rcvr;
	int imhash, followptr;

	t = methtype(t0);

	// if this is too slow,
	// could sort these first
	// and then do one loop.

	// could also do full type compare
	// instead of using hash, but have to
	// avoid checking receivers, and
	// typehash already does that for us.
	// also, it&#39;s what the runtime will do,
	// so we can both be wrong together.

	for(im=iface-&gt;type; im; im=im-&gt;down) {
		imhash = typehash(im-&gt;type);
		tm = ifacelookdot(im-&gt;sym, t, &amp;followptr);
		if(tm == T || typehash(tm-&gt;type) != imhash) {
			*m = im;
			*samename = tm;
			return 0;
		}
		// if pointer receiver in method,
		// the method does not exist for value types.
		rcvr = getthisx(tm-&gt;type)-&gt;type-&gt;type;
		if(isptr[rcvr-&gt;etype] &amp;&amp; !isptr[t0-&gt;etype] &amp;&amp; !followptr &amp;&amp; !isifacemethod(tm)) {
			if(debug[&#39;r&#39;])
				yyerror(&#34;interface pointer mismatch&#34;);
			*m = im;
			*samename = nil;
			return 0;
		}
	}
	return 1;
}

// check whether interface type i1 satisifes interface type i2.
int
ifaceokI2I(Type *i1, Type *i2, Type **m)
{
	Type *m1, *m2;

	// if this is too slow,
	// could sort these first
	// and then do one loop.

	for(m2=i2-&gt;type; m2; m2=m2-&gt;down) {
		for(m1=i1-&gt;type; m1; m1=m1-&gt;down)
			if(m1-&gt;sym == m2-&gt;sym &amp;&amp; typehash(m1) == typehash(m2))
				goto found;
		*m = m2;
		return 0;
	found:;
	}
	return 1;
}

void
runifacechecks(void)
{
	Icheck *p;
	int lno, wrong, needexplicit;
	Type *m, *t, *iface, *samename;

	lno = lineno;
	for(p=icheck; p; p=p-&gt;next) {
		lineno = p-&gt;lineno;
		wrong = 0;
		needexplicit = 0;
		m = nil;
		samename = nil;
		if(isinter(p-&gt;dst) &amp;&amp; isinter(p-&gt;src)) {
			iface = p-&gt;dst;
			t = p-&gt;src;
			needexplicit = !ifaceokI2I(t, iface, &amp;m);
		}
		else if(isinter(p-&gt;dst)) {
			t = p-&gt;src;
			iface = p-&gt;dst;
			wrong = !ifaceokT2I(t, iface, &amp;m, &amp;samename);
		} else {
			t = p-&gt;dst;
			iface = p-&gt;src;
			wrong = !ifaceokT2I(t, iface, &amp;m, &amp;samename);
			needexplicit = 1;
		}
		if(wrong) {
			if(samename)
				yyerror(&#34;%T is not %T\n\tmissing %S%hhT\n\tdo have %S%hhT&#34;,
					t, iface, m-&gt;sym, m-&gt;type, samename-&gt;sym, samename-&gt;type);
			else
				yyerror(&#34;%T is not %T\n\tmissing %S%hhT&#34;, t, iface, m-&gt;sym, m-&gt;type);
		}
		else if(!p-&gt;explicit &amp;&amp; needexplicit) {
			if(m) {
				if(samename)
					yyerror(&#34;need type assertion to use %T as %T\n\tmissing %S %hhT\n\tdo have %S%hhT&#34;,
						p-&gt;src, p-&gt;dst, m-&gt;sym, m-&gt;type, samename-&gt;sym, samename-&gt;type);
				else
					yyerror(&#34;need type assertion to use %T as %T\n\tmissing %S%hhT&#34;,
						p-&gt;src, p-&gt;dst, m-&gt;sym, m-&gt;type);
			} else
				yyerror(&#34;need type assertion to use %T as %T&#34;,
					p-&gt;src, p-&gt;dst);
		}
	}
	lineno = lno;
}

/*
 * even simpler simtype; get rid of ptr, bool.
 * assuming that the front end has rejected
 * all the invalid conversions (like ptr -&gt; bool)
 */
int
simsimtype(Type *t)
{
	int et;

	if(t == 0)
		return 0;

	et = simtype[t-&gt;etype];
	switch(et) {
	case TPTR32:
		et = TUINT32;
		break;
	case TPTR64:
		et = TUINT64;
		break;
	case TBOOL:
		et = TUINT8;
		break;
	}
	return et;
}

NodeList*
concat(NodeList *a, NodeList *b)
{
	if(a == nil)
		return b;
	if(b == nil)
		return a;

	a-&gt;end-&gt;next = b;
	a-&gt;end = b-&gt;end;
	b-&gt;end = nil;
	return a;
}

NodeList*
list1(Node *n)
{
	NodeList *l;

	if(n == nil)
		return nil;
	if(n-&gt;op == OBLOCK &amp;&amp; n-&gt;ninit == nil)
		return n-&gt;list;
	l = mal(sizeof *l);
	l-&gt;n = n;
	l-&gt;end = l;
	return l;
}

NodeList*
list(NodeList *l, Node *n)
{
	return concat(l, list1(n));
}

NodeList*
listtreecopy(NodeList *l)
{
	NodeList *out;

	out = nil;
	for(; l; l=l-&gt;next)
		out = list(out, treecopy(l-&gt;n));
	return out;
}

Node*
liststmt(NodeList *l)
{
	Node *n;

	n = nod(OBLOCK, N, N);
	n-&gt;list = l;
	if(l)
		n-&gt;lineno = l-&gt;n-&gt;lineno;
	return n;
}

/*
 * return nelem of list
 */
int
count(NodeList *l)
{
	int n;

	n = 0;
	for(; l; l=l-&gt;next)
		n++;
	return n;
}

/*
 * return nelem of list
 */
int
structcount(Type *t)
{
	int v;
	Iter s;

	v = 0;
	for(t = structfirst(&amp;s, &amp;t); t != T; t = structnext(&amp;s))
		v++;
	return v;
}

/*
 * return power of 2 of the constant
 * operand. -1 if it is not a power of 2.
 * 1000+ if it is a -(power of 2)
 */
int
powtwo(Node *n)
{
	uvlong v, b;
	int i;

	if(n == N || n-&gt;op != OLITERAL || n-&gt;type == T)
		goto no;
	if(!isint[n-&gt;type-&gt;etype])
		goto no;

	v = mpgetfix(n-&gt;val.u.xval);
	b = 1ULL;
	for(i=0; i&lt;64; i++) {
		if(b == v)
			return i;
		b = b&lt;&lt;1;
	}

	if(!issigned[n-&gt;type-&gt;etype])
		goto no;

	v = -v;
	b = 1ULL;
	for(i=0; i&lt;64; i++) {
		if(b == v)
			return i+1000;
		b = b&lt;&lt;1;
	}

no:
	return -1;
}

/*
 * return the unsigned type for
 * a signed integer type.
 * returns T if input is not a
 * signed integer type.
 */
Type*
tounsigned(Type *t)
{

	// this is types[et+1], but not sure
	// that this relation is immutable
	switch(t-&gt;etype) {
	default:
		print(&#34;tounsigned: unknown type %T\n&#34;, t);
		t = T;
		break;
	case TINT:
		t = types[TUINT];
		break;
	case TINT8:
		t = types[TUINT8];
		break;
	case TINT16:
		t = types[TUINT16];
		break;
	case TINT32:
		t = types[TUINT32];
		break;
	case TINT64:
		t = types[TUINT64];
		break;
	}
	return t;
}

/*
 * magic number for signed division
 * see hacker&#39;s delight chapter 10
 */
void
smagic(Magic *m)
{
	int p;
	uint64 ad, anc, delta, q1, r1, q2, r2, t;
	uint64 mask, two31;

	m-&gt;bad = 0;
	switch(m-&gt;w) {
	default:
		m-&gt;bad = 1;
		return;
	case 8:
		mask = 0xffLL;
		break;
	case 16:
		mask = 0xffffLL;
		break;
	case 32:
		mask = 0xffffffffLL;
		break;
	case 64:
		mask = 0xffffffffffffffffLL;
		break;
	}
	two31 = mask ^ (mask&gt;&gt;1);

	p = m-&gt;w-1;
	ad = m-&gt;sd;
	if(m-&gt;sd &lt; 0)
		ad = -m-&gt;sd;

	// bad denominators
	if(ad == 0 || ad == 1 || ad == two31) {
		m-&gt;bad = 1;
		return;
	}

	t = two31;
	ad &amp;= mask;

	anc = t - 1 - t%ad;
	anc &amp;= mask;

	q1 = two31/anc;
	r1 = two31 - q1*anc;
	q1 &amp;= mask;
	r1 &amp;= mask;

	q2 = two31/ad;
	r2 = two31 - q2*ad;
	q2 &amp;= mask;
	r2 &amp;= mask;

	for(;;) {
		p++;
		q1 &lt;&lt;= 1;
		r1 &lt;&lt;= 1;
		q1 &amp;= mask;
		r1 &amp;= mask;
		if(r1 &gt;= anc) {
			q1++;
			r1 -= anc;
			q1 &amp;= mask;
			r1 &amp;= mask;
		}

		q2 &lt;&lt;= 1;
		r2 &lt;&lt;= 1;
		q2 &amp;= mask;
		r2 &amp;= mask;
		if(r2 &gt;= ad) {
			q2++;
			r2 -= ad;
			q2 &amp;= mask;
			r2 &amp;= mask;
		}

		delta = ad - r2;
		delta &amp;= mask;
		if(q1 &lt; delta || (q1 == delta &amp;&amp; r1 == 0)) {
			continue;
		}
		break;
	}

	m-&gt;sm = q2+1;
	if(m-&gt;sm &amp; two31)
		m-&gt;sm |= ~mask;
	m-&gt;s = p-m-&gt;w;
}

/*
 * magic number for unsigned division
 * see hacker&#39;s delight chapter 10
 */
void
umagic(Magic *m)
{
	int p;
	uint64 nc, delta, q1, r1, q2, r2;
	uint64 mask, two31;

	m-&gt;bad = 0;
	m-&gt;ua = 0;

	switch(m-&gt;w) {
	default:
		m-&gt;bad = 1;
		return;
	case 8:
		mask = 0xffLL;
		break;
	case 16:
		mask = 0xffffLL;
		break;
	case 32:
		mask = 0xffffffffLL;
		break;
	case 64:
		mask = 0xffffffffffffffffLL;
		break;
	}
	two31 = mask ^ (mask&gt;&gt;1);

	m-&gt;ud &amp;= mask;
	if(m-&gt;ud == 0 || m-&gt;ud == two31) {
		m-&gt;bad = 1;
		return;
	}
	nc = mask - (-m-&gt;ud&amp;mask)%m-&gt;ud;
	p = m-&gt;w-1;

	q1 = two31/nc;
	r1 = two31 - q1*nc;
	q1 &amp;= mask;
	r1 &amp;= mask;

	q2 = (two31-1) / m-&gt;ud;
	r2 = (two31-1) - q2*m-&gt;ud;
	q2 &amp;= mask;
	r2 &amp;= mask;

	for(;;) {
		p++;
		if(r1 &gt;= nc-r1) {
			q1 &lt;&lt;= 1;
			q1++;
			r1 &lt;&lt;= 1;
			r1 -= nc;
		} else {
			q1 &lt;&lt;= 1;
			r1 &lt;&lt;= 1;
		}
		q1 &amp;= mask;
		r1 &amp;= mask;
		if(r2+1 &gt;= m-&gt;ud-r2) {
			if(q2 &gt;= two31-1) {
				m-&gt;ua = 1;
			}
			q2 &lt;&lt;= 1;
			q2++;
			r2 &lt;&lt;= 1;
			r2++;
			r2 -= m-&gt;ud;
		} else {
			if(q2 &gt;= two31) {
				m-&gt;ua = 1;
			}
			q2 &lt;&lt;= 1;
			r2 &lt;&lt;= 1;
			r2++;
		}
		q2 &amp;= mask;
		r2 &amp;= mask;

		delta = m-&gt;ud - 1 - r2;
		delta &amp;= mask;

		if(p &lt; m-&gt;w+m-&gt;w)
		if(q1 &lt; delta || (q1 == delta &amp;&amp; r1 == 0)) {
			continue;
		}
		break;
	}
	m-&gt;um = q2+1;
	m-&gt;s = p-m-&gt;w;
}

Sym*
ngotype(Node *n)
{
	if(n-&gt;sym != S &amp;&amp; strncmp(n-&gt;sym-&gt;name, &#34;autotmp_&#34;, 8) != 0)
	if(n-&gt;type-&gt;etype != TFUNC || n-&gt;type-&gt;thistuple == 0)
	if(n-&gt;type-&gt;etype != TSTRUCT || n-&gt;type-&gt;funarg == 0)
		return typename(n-&gt;type)-&gt;left-&gt;sym;
	return S;
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
