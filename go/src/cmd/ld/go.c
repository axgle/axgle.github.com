<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/ld/go.c</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/ld/go.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go-specific code shared across loaders (5l, 6l, 8l).

#include	&#34;l.h&#34;
#include	&#34;../ld/lib.h&#34;

// accumulate all type information from .6 files.
// check for inconsistencies.

// TODO:
//	generate debugging section in binary.
//	once the dust settles, try to move some code to
//		libmach, so that other linkers and ar can share.

/*
 *	package import data
 */
typedef struct Import Import;
struct Import
{
	Import *hash;	// next in hash table
	char *prefix;	// &#34;type&#34;, &#34;var&#34;, &#34;func&#34;, &#34;const&#34;
	char *name;
	char *def;
	char *file;
};
enum {
	NIHASH = 1024
};
static Import *ihash[NIHASH];
static int nimport;

static int
hashstr(char *name)
{
	int h;
	char *cp;

	h = 0;
	for(cp = name; *cp; h += *cp++)
		h *= 1119;
	// not if(h &lt; 0) h = ~h, because gcc 4.3 -O2 miscompiles it.
	h &amp;= 0xffffff;
	return h;
}

static Import *
ilookup(char *name)
{
	int h;
	Import *x;

	h = hashstr(name) % NIHASH;
	for(x=ihash[h]; x; x=x-&gt;hash)
		if(x-&gt;name[0] == name[0] &amp;&amp; strcmp(x-&gt;name, name) == 0)
			return x;
	x = mal(sizeof *x);
	x-&gt;name = strdup(name);
	x-&gt;hash = ihash[h];
	ihash[h] = x;
	nimport++;
	return x;
}

static void loadpkgdata(char*, char*, int);
static void loaddynld(char*, char*, int);
static int parsemethod(char**, char*, char**);
static int parsepkgdata(char*, char**, char*, char**, char**, char**);

void
ldpkg(Biobuf *f, int64 len, char *filename)
{
	char *data, *p0, *p1;

	if(debug[&#39;g&#39;])
		return;

	if((int)len != len) {
		fprint(2, &#34;%s: too much pkg data in %s\n&#34;, argv0, filename);
		return;
	}
	data = mal(len+1);
	if(Bread(f, data, len) != len) {
		fprint(2, &#34;%s: short pkg read %s\n&#34;, argv0, filename);
		return;
	}
	data[len] = &#39;\0&#39;;

	// first \n$$ marks beginning of exports - skip rest of line
	p0 = strstr(data, &#34;\n$$&#34;);
	if(p0 == nil)
		return;
	p0 += 3;
	while(*p0 != &#39;\n&#39; &amp;&amp; *p0 != &#39;\0&#39;)
		p0++;

	// second marks end of exports / beginning of local data
	p1 = strstr(p0, &#34;\n$$&#34;);
	if(p1 == nil) {
		fprint(2, &#34;%s: cannot find end of exports in %s\n&#34;, argv0, filename);
		return;
	}
	while(p0 &lt; p1 &amp;&amp; (*p0 == &#39; &#39; || *p0 == &#39;\t&#39; || *p0 == &#39;\n&#39;))
		p0++;
	if(p0 &lt; p1) {
		if(strncmp(p0, &#34;package &#34;, 8) != 0) {
			fprint(2, &#34;%s: bad package section in %s - %s\n&#34;, argv0, filename, p0);
			return;
		}
		p0 += 8;
		while(*p0 == &#39; &#39; || *p0 == &#39;\t&#39; || *p0 == &#39;\n&#39;)
			p0++;
		while(*p0 != &#39; &#39; &amp;&amp; *p0 != &#39;\t&#39; &amp;&amp; *p0 != &#39;\n&#39;)
			p0++;

		loadpkgdata(filename, p0, p1 - p0);
	}

	// local types begin where exports end.
	// skip rest of line after $$ we found above
	p0 = p1 + 3;
	while(*p0 != &#39;\n&#39; &amp;&amp; *p0 != &#39;\0&#39;)
		p0++;

	// local types end at next \n$$.
	p1 = strstr(p0, &#34;\n$$&#34;);
	if(p1 == nil) {
		fprint(2, &#34;%s: cannot find end of local types in %s\n&#34;, argv0, filename);
		return;
	}

	loadpkgdata(filename, p0, p1 - p0);

	// look for dynld section
	p0 = strstr(p1, &#34;\n$$  // dynld&#34;);
	if(p0 != nil) {
		p0 = strchr(p0+1, &#39;\n&#39;);
		if(p0 == nil) {
			fprint(2, &#34;%s: found $$ // dynld but no newline in %s\n&#34;, argv0, filename);
			return;
		}
		p1 = strstr(p0, &#34;\n$$&#34;);
		if(p1 == nil)
			p1 = strstr(p0, &#34;\n!\n&#34;);
		if(p1 == nil) {
			fprint(2, &#34;%s: cannot find end of // dynld section in %s\n&#34;, argv0, filename);
			return;
		}
		loaddynld(filename, p0 + 1, p1 - p0);
	}
}

/*
 * a and b don&#39;t match.
 * is one a forward declaration and the other a valid completion?
 * if so, return the one to keep.
 */
char*
forwardfix(char *a, char *b)
{
	char *t;

	if(strlen(a) &gt; strlen(b)) {
		t = a;
		a = b;
		b = t;
	}
	if(strcmp(a, &#34;struct&#34;) == 0 &amp;&amp; strncmp(b, &#34;struct &#34;, 7) == 0)
		return b;
	if(strcmp(a, &#34;interface&#34;) == 0 &amp;&amp; strncmp(b, &#34;interface &#34;, 10) == 0)
		return b;
	return nil;
}

static void
loadpkgdata(char *file, char *data, int len)
{
	char *p, *ep, *prefix, *name, *def, *ndef;
	Import *x;

	file = strdup(file);
	p = data;
	ep = data + len;
	while(parsepkgdata(file, &amp;p, ep, &amp;prefix, &amp;name, &amp;def) &gt; 0) {
		x = ilookup(name);
		if(x-&gt;prefix == nil) {
			x-&gt;prefix = prefix;
			x-&gt;def = def;
			x-&gt;file = file;
		} else if(strcmp(x-&gt;prefix, prefix) != 0) {
			fprint(2, &#34;%s: conflicting definitions for %s\n&#34;, argv0, name);
			fprint(2, &#34;%s:\t%s %s ...\n&#34;, x-&gt;file, x-&gt;prefix, name);
			fprint(2, &#34;%s:\t%s %s ...\n&#34;, file, prefix, name);
			nerrors++;
		} else if(strcmp(x-&gt;def, def) == 0) {
			// fine
		} else if((ndef = forwardfix(x-&gt;def, def)) != nil) {
			x-&gt;def = ndef;
		} else {
			fprint(2, &#34;%s: conflicting definitions for %s\n&#34;, argv0, name);
			fprint(2, &#34;%s:\t%s %s %s\n&#34;, x-&gt;file, x-&gt;prefix, name, x-&gt;def);
			fprint(2, &#34;%s:\t%s %s %s\n&#34;, file, prefix, name, def);
			nerrors++;
		}
	}
}

static int
parsepkgdata(char *file, char **pp, char *ep, char **prefixp, char **namep, char **defp)
{
	char *p, *prefix, *name, *def, *edef, *meth;
	int n;

	// skip white space
	p = *pp;
	while(p &lt; ep &amp;&amp; (*p == &#39; &#39; || *p == &#39;\t&#39; || *p == &#39;\n&#39;))
		p++;
	if(p == ep || strncmp(p, &#34;$$\n&#34;, 3) == 0)
		return 0;

	// prefix: (var|type|func|const)
	prefix = p;
	if(p + 6 &gt; ep)
		return -1;
	if(strncmp(p, &#34;var &#34;, 4) == 0)
		p += 4;
	else if(strncmp(p, &#34;type &#34;, 5) == 0)
		p += 5;
	else if(strncmp(p, &#34;func &#34;, 5) == 0)
		p += 5;
	else if(strncmp(p, &#34;const &#34;, 6) == 0)
		p += 6;
	else {
		fprint(2, &#34;%s: confused in pkg data near &lt;&lt;%.40s&gt;&gt;\n&#34;, argv0, prefix);
		nerrors++;
		return -1;
	}
	p[-1] = &#39;\0&#39;;

	// name: a.b followed by space
	name = p;
	while(p &lt; ep &amp;&amp; *p != &#39; &#39;)
		p++;
	if(p &gt;= ep)
		return -1;
	*p++ = &#39;\0&#39;;

	// def: free form to new line
	def = p;
	while(p &lt; ep &amp;&amp; *p != &#39;\n&#39;)
		p++;
	if(p &gt;= ep)
		return -1;
	edef = p;
	*p++ = &#39;\0&#39;;

	// include methods on successive lines in def of named type
	while(parsemethod(&amp;p, ep, &amp;meth) &gt; 0) {
		*edef++ = &#39;\n&#39;;	// overwrites &#39;\0&#39;
		if(edef+1 &gt; meth) {
			// We want to indent methods with a single \t.
			// 6g puts at least one char of indent before all method defs,
			// so there will be room for the \t.  If the method def wasn&#39;t
			// indented we could do something more complicated,
			// but for now just diagnose the problem and assume
			// 6g will keep indenting for us.
			fprint(2, &#34;%s: %s: expected methods to be indented %p %p %.10s\n&#34;, argv0,
				file, edef, meth, meth);
			nerrors++;
			return -1;
		}
		*edef++ = &#39;\t&#39;;
		n = strlen(meth);
		memmove(edef, meth, n);
		edef += n;
	}

	// done
	*pp = p;
	*prefixp = prefix;
	*namep = name;
	*defp = def;
	return 1;
}

static int
parsemethod(char **pp, char *ep, char **methp)
{
	char *p;

	// skip white space
	p = *pp;
	while(p &lt; ep &amp;&amp; (*p == &#39; &#39; || *p == &#39;\t&#39;))
		p++;
	if(p == ep)
		return 0;

	// if it says &#34;func (&#34;, it&#39;s a method
	if(p + 6 &gt;= ep || strncmp(p, &#34;func (&#34;, 6) != 0)
		return 0;

	// definition to end of line
	*methp = p;
	while(p &lt; ep &amp;&amp; *p != &#39;\n&#39;)
		p++;
	if(p &gt;= ep) {
		fprint(2, &#34;%s: lost end of line in method definition\n&#34;, argv0);
		*pp = ep;
		return -1;
	}
	*p++ = &#39;\0&#39;;
	*pp = p;
	return 1;
}

static void
loaddynld(char *file, char *p, int n)
{
	char *next, *name, *def, *p0, *lib;
	Sym *s;

	p[n] = &#39;\0&#39;;

	p0 = p;
	for(; *p; p=next) {
		next = strchr(p, &#39;\n&#39;);
		if(next == nil)
			next = &#34;&#34;;
		else
			*next++ = &#39;\0&#39;;
		p0 = p;
		if(strncmp(p, &#34;dynld &#34;, 6) != 0)
			goto err;
		p += 6;
		name = p;
		p = strchr(name, &#39; &#39;);
		if(p == nil)
			goto err;
		while(*p == &#39; &#39;)
			p++;
		def = p;
		p = strchr(def, &#39; &#39;);
		if(p == nil)
			goto err;
		while(*p == &#39; &#39;)
			p++;
		lib = p;

		// successful parse: now can edit the line
		*strchr(name, &#39; &#39;) = 0;
		*strchr(def, &#39; &#39;) = 0;

		s = lookup(name, 0);
		s-&gt;dynldlib = lib;
		s-&gt;dynldname = def;
	}
	return;

err:
	fprint(2, &#34;%s: invalid dynld line: %s\n&#34;, argv0, p0);
	nerrors++;
}

static void mark(Sym*);
static int markdepth;

static void
markdata(Prog *p, Sym *s)
{
	markdepth++;
	if(p != P &amp;&amp; debug[&#39;v&#39;] &gt; 1)
		Bprint(&amp;bso, &#34;%d markdata %s\n&#34;, markdepth, s-&gt;name);
	for(; p != P; p=p-&gt;dlink)
		if(p-&gt;to.sym)
			mark(p-&gt;to.sym);
	markdepth--;
}

static void
marktext(Prog *p)
{
	Auto *a;

	if(p == P)
		return;
	if(p-&gt;as != ATEXT) {
		diag(&#34;marktext: %P&#34;, p);
		return;
	}
	for(a=p-&gt;to.autom; a; a=a-&gt;link)
		mark(a-&gt;gotype);
	markdepth++;
	if(debug[&#39;v&#39;] &gt; 1)
		Bprint(&amp;bso, &#34;%d marktext %s\n&#34;, markdepth, p-&gt;from.sym-&gt;name);
	for(a=p-&gt;to.autom; a; a=a-&gt;link)
		mark(a-&gt;gotype);
	for(p=p-&gt;link; p != P; p=p-&gt;link) {
		if(p-&gt;as == ATEXT || p-&gt;as == ADATA || p-&gt;as == AGLOBL)
			break;
		if(p-&gt;from.sym)
			mark(p-&gt;from.sym);
		if(p-&gt;to.sym)
			mark(p-&gt;to.sym);
	}
	markdepth--;
}

static void
mark(Sym *s)
{
	if(s == S || s-&gt;reachable)
		return;
	s-&gt;reachable = 1;
	if(s-&gt;text)
		marktext(s-&gt;text);
	if(s-&gt;data)
		markdata(s-&gt;data, s);
	if(s-&gt;gotype)
		mark(s-&gt;gotype);
}

static void
sweeplist(Prog **first, Prog **last)
{
	int reachable;
	Prog *p, *q;

	reachable = 1;
	q = P;
	for(p=*first; p != P; p=p-&gt;link) {
		switch(p-&gt;as) {
		case ATEXT:
		case ADATA:
		case AGLOBL:
			reachable = p-&gt;from.sym-&gt;reachable;
		}
		if(reachable) {
			if(q == P)
				*first = p;
			else
				q-&gt;link = p;
			q = p;
		}
	}
	if(q == P)
		*first = P;
	else
		q-&gt;link = P;
	*last = q;
}

static char*
morename[] =
{
	&#34;runtime·morestack&#34;,
	&#34;runtime·morestackx&#34;,

	&#34;runtime·morestack00&#34;,
	&#34;runtime·morestack10&#34;,
	&#34;runtime·morestack01&#34;,
	&#34;runtime·morestack11&#34;,

	&#34;runtime·morestack8&#34;,
	&#34;runtime·morestack16&#34;,
	&#34;runtime·morestack24&#34;,
	&#34;runtime·morestack32&#34;,
	&#34;runtime·morestack40&#34;,
	&#34;runtime·morestack48&#34;,
};

void
deadcode(void)
{
	int i;

	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f deadcode\n&#34;, cputime());

	mark(lookup(INITENTRY, 0));
	for(i=0; i&lt;nelem(morename); i++)
		mark(lookup(morename[i], 0));

	// remove dead data
	sweeplist(&amp;datap, &amp;edatap);
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
