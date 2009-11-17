<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cc/dpchk.c</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/cc/dpchk.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/cc/dpchk.c
// http://code.google.com/p/inferno-os/source/browse/utils/cc/dpchk.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the &#34;Software&#34;), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

#include	&#34;cc.h&#34;
#include	&#34;y.tab.h&#34;

enum
{
	Fnone	= 0,
	Fl,
	Fvl,
	Fignor,
	Fstar,
	Fadj,

	Fverb	= 10,
};

typedef	struct	Tprot	Tprot;
struct	Tprot
{
	Type*	type;
	Bits	flag;
	Tprot*	link;
};

typedef	struct	Tname	Tname;
struct	Tname
{
	char*	name;
	int	param;
	Tname*	link;
};

static	Type*	indchar;
static	uchar	flagbits[512];
static	char	fmtbuf[100];
static	int	lastadj;
static	int	lastverb;
static	int	nstar;
static	Tprot*	tprot;
static	Tname*	tname;

void
argflag(int c, int v)
{

	switch(v) {
	case Fignor:
	case Fstar:
	case Fl:
	case Fvl:
		flagbits[c] = v;
		break;
	case Fverb:
		flagbits[c] = lastverb;
/*print(&#34;flag-v %c %d\n&#34;, c, lastadj);*/
		lastverb++;
		break;
	case Fadj:
		flagbits[c] = lastadj;
/*print(&#34;flag-l %c %d\n&#34;, c, lastadj);*/
		lastadj++;
		break;
	}
}

Bits
getflag(char *s)
{
	Bits flag;
	int f;
	char *fmt;
	Rune c;

	fmt = fmtbuf;
	flag = zbits;
	nstar = 0;
	for(;;) {
		s += chartorune(&amp;c, s);
		if(c == 0 || c &gt;= nelem(flagbits))
			break;
		fmt += runetochar(fmt, &amp;c);
		f = flagbits[c];
		switch(f) {
		case Fnone:
			argflag(c, Fverb);
			f = flagbits[c];
			break;
		case Fstar:
			nstar++;
		case Fignor:
			continue;
		case Fl:
			if(bset(flag, Fl))
				flag = bor(flag, blsh(Fvl));
		}
		flag = bor(flag, blsh(f));
		if(f &gt;= Fverb)
			break;
	}
	*fmt = 0;
	return flag;
}

void
newprot(Sym *m, Type *t, char *s)
{
	Bits flag;
	Tprot *l;

	if(t == T) {
		warn(Z, &#34;%s: newprot: type not defined&#34;, m-&gt;name);
		return;
	}
	flag = getflag(s);
	for(l=tprot; l; l=l-&gt;link)
		if(beq(flag, l-&gt;flag) &amp;&amp; sametype(t, l-&gt;type))
			return;
	l = alloc(sizeof(*l));
	l-&gt;type = t;
	l-&gt;flag = flag;
	l-&gt;link = tprot;
	tprot = l;
}

void
newname(char *s, int p)
{
	Tname *l;

	for(l=tname; l; l=l-&gt;link)
		if(strcmp(l-&gt;name, s) == 0) {
			if(l-&gt;param != p)
				yyerror(&#34;vargck %s already defined\n&#34;, s);
			return;
		}
	l = alloc(sizeof(*l));
	l-&gt;name = s;
	l-&gt;param = p;
	l-&gt;link = tname;
	tname = l;
}

void
arginit(void)
{
	int i;

/* debug[&#39;F&#39;] = 1;*/
/* debug[&#39;w&#39;] = 1;*/

	lastadj = Fadj;
	lastverb = Fverb;
	indchar = typ(TIND, types[TCHAR]);

	memset(flagbits, Fnone, sizeof(flagbits));

	for(i=&#39;0&#39;; i&lt;=&#39;9&#39;; i++)
		argflag(i, Fignor);
	argflag(&#39;.&#39;, Fignor);
	argflag(&#39;#&#39;, Fignor);
	argflag(&#39;u&#39;, Fignor);
	argflag(&#39;h&#39;, Fignor);
	argflag(&#39;+&#39;, Fignor);
	argflag(&#39;-&#39;, Fignor);

	argflag(&#39;*&#39;, Fstar);
	argflag(&#39;l&#39;, Fl);

	argflag(&#39;o&#39;, Fverb);
	flagbits[&#39;x&#39;] = flagbits[&#39;o&#39;];
	flagbits[&#39;X&#39;] = flagbits[&#39;o&#39;];
}

static char*
getquoted(void)
{
	int c;
	char *t;
	Rune r;

	c = getnsc();
	if(c != &#39;&#34;&#39;)
		return nil;
	t = fmtbuf;
	for(;;) {
		r = getr();
		if(r == &#39; &#39; || r == &#39;\n&#39;)
			return nil;
		if(r == &#39;&#34;&#39;)
			break;
		t += runetochar(t, &amp;r);
	}
	*t = 0;
	return strdup(fmtbuf);
}

void
pragvararg(void)
{
	Sym *s;
	int n, c;
	char *t;
	Type *ty;

	if(!debug[&#39;F&#39;])
		goto out;
	s = getsym();
	if(s &amp;&amp; strcmp(s-&gt;name, &#34;argpos&#34;) == 0)
		goto ckpos;
	if(s &amp;&amp; strcmp(s-&gt;name, &#34;type&#34;) == 0)
		goto cktype;
	if(s &amp;&amp; strcmp(s-&gt;name, &#34;flag&#34;) == 0)
		goto ckflag;
	yyerror(&#34;syntax in #pragma varargck&#34;);
	goto out;

ckpos:
/*#pragma	varargck	argpos	warn	2*/
	s = getsym();
	if(s == S)
		goto bad;
	n = getnsn();
	if(n &lt; 0)
		goto bad;
	newname(s-&gt;name, n);
	goto out;

ckflag:
/*#pragma	varargck	flag	&#39;c&#39;*/
	c = getnsc();
	if(c != &#39;\&#39;&#39;)
		goto bad;
	c = getr();
	if(c == &#39;\\&#39;)
		c = getr();
	else if(c == &#39;\&#39;&#39;)
		goto bad;
	if(c == &#39;\n&#39;)
		goto bad;
	if(getc() != &#39;\&#39;&#39;)
		goto bad;
	argflag(c, Fignor);
	goto out;

cktype:
/*#pragma	varargck	type	O	int*/
	t = getquoted();
	if(t == nil)
		goto bad;
	s = getsym();
	if(s == S)
		goto bad;
	ty = s-&gt;type;
	while((c = getnsc()) == &#39;*&#39;)
		ty = typ(TIND, ty);
	unget(c);
	newprot(s, ty, t);
	goto out;

bad:
	yyerror(&#34;syntax in #pragma varargck&#34;);

out:
	while(getnsc() != &#39;\n&#39;)
		;
}

Node*
nextarg(Node *n, Node **a)
{
	if(n == Z) {
		*a = Z;
		return Z;
	}
	if(n-&gt;op == OLIST) {
		*a = n-&gt;left;
		return n-&gt;right;
	}
	*a = n;
	return Z;
}

void
checkargs(Node *nn, char *s, int pos)
{
	Node *a, *n;
	Bits flag;
	Tprot *l;

	if(!debug[&#39;F&#39;])
		return;
	n = nn;
	for(;;) {
		s = strchr(s, &#39;%&#39;);
		if(s == 0) {
			nextarg(n, &amp;a);
			if(a != Z)
				warn(nn, &#34;more arguments than format %T&#34;,
					a-&gt;type);
			return;
		}
		s++;
		flag = getflag(s);
		while(nstar &gt; 0) {
			n = nextarg(n, &amp;a);
			pos++;
			nstar--;
			if(a == Z) {
				warn(nn, &#34;more format than arguments %s&#34;,
					fmtbuf);
				return;
			}
			if(a-&gt;type == T)
				continue;
			if(!sametype(types[TINT], a-&gt;type) &amp;&amp;
			   !sametype(types[TUINT], a-&gt;type))
				warn(nn, &#34;format mismatch &#39;*&#39; in %s %T, arg %d&#34;,
					fmtbuf, a-&gt;type, pos);
		}
		for(l=tprot; l; l=l-&gt;link)
			if(sametype(types[TVOID], l-&gt;type)) {
				if(beq(flag, l-&gt;flag)) {
					s++;
					goto loop;
				}
			}

		n = nextarg(n, &amp;a);
		pos++;
		if(a == Z) {
			warn(nn, &#34;more format than arguments %s&#34;,
				fmtbuf);
			return;
		}
		if(a-&gt;type == 0)
			continue;
		for(l=tprot; l; l=l-&gt;link)
			if(sametype(a-&gt;type, l-&gt;type)) {
/*print(&#34;checking %T/%ulx %T/%ulx\n&#34;, a-&gt;type, flag.b[0], l-&gt;type, l-&gt;flag.b[0]);*/
				if(beq(flag, l-&gt;flag))
					goto loop;
			}
		warn(nn, &#34;format mismatch %s %T, arg %d&#34;, fmtbuf, a-&gt;type, pos);
	loop:;
	}
}

void
dpcheck(Node *n)
{
	char *s;
	Node *a, *b;
	Tname *l;
	int i;

	if(n == Z)
		return;
	b = n-&gt;left;
	if(b == Z || b-&gt;op != ONAME)
		return;
	s = b-&gt;sym-&gt;name;
	for(l=tname; l; l=l-&gt;link)
		if(strcmp(s, l-&gt;name) == 0)
			break;
	if(l == 0)
		return;

	i = l-&gt;param;
	b = n-&gt;right;
	while(i &gt; 0) {
		b = nextarg(b, &amp;a);
		i--;
	}
	if(a == Z) {
		warn(n, &#34;cant find format arg&#34;);
		return;
	}
	if(!sametype(indchar, a-&gt;type)) {
		warn(n, &#34;format arg type %T&#34;, a-&gt;type);
		return;
	}
	if(a-&gt;op != OADDR || a-&gt;left-&gt;op != ONAME || a-&gt;left-&gt;sym != symstring) {
/*		warn(n, &#34;format arg not constant string&#34;);*/
		return;
	}
	s = a-&gt;left-&gt;cstring;
	checkargs(b, s, l-&gt;param);
}

void
pragpack(void)
{
	Sym *s;

	packflg = 0;
	s = getsym();
	if(s) {
		packflg = atoi(s-&gt;name+1);
		if(strcmp(s-&gt;name, &#34;on&#34;) == 0 ||
		   strcmp(s-&gt;name, &#34;yes&#34;) == 0)
			packflg = 1;
	}
	while(getnsc() != &#39;\n&#39;)
		;
	if(debug[&#39;f&#39;])
		if(packflg)
			print(&#34;%4ld: pack %d\n&#34;, lineno, packflg);
		else
			print(&#34;%4ld: pack off\n&#34;, lineno);
}

void
pragfpround(void)
{
	Sym *s;

	fproundflg = 0;
	s = getsym();
	if(s) {
		fproundflg = atoi(s-&gt;name+1);
		if(strcmp(s-&gt;name, &#34;on&#34;) == 0 ||
		   strcmp(s-&gt;name, &#34;yes&#34;) == 0)
			fproundflg = 1;
	}
	while(getnsc() != &#39;\n&#39;)
		;
	if(debug[&#39;f&#39;])
		if(fproundflg)
			print(&#34;%4ld: fproundflg %d\n&#34;, lineno, fproundflg);
		else
			print(&#34;%4ld: fproundflg off\n&#34;, lineno);
}

void
pragtextflag(void)
{
	Sym *s;

	textflag = 0;
	s = getsym();
	textflag = 7;
	if(s)
		textflag = atoi(s-&gt;name+1);
	while(getnsc() != &#39;\n&#39;)
		;
	if(debug[&#39;f&#39;])
		print(&#34;%4ld: textflag %d\n&#34;, lineno, textflag);
}

void
pragincomplete(void)
{
	Sym *s;
	Type *t;
	int istag, w, et;

	istag = 0;
	s = getsym();
	if(s == nil)
		goto out;
	et = 0;
	w = s-&gt;lexical;
	if(w == LSTRUCT)
		et = TSTRUCT;
	else if(w == LUNION)
		et = TUNION;
	if(et != 0){
		s = getsym();
		if(s == nil){
			yyerror(&#34;missing struct/union tag in pragma incomplete&#34;);
			goto out;
		}
		if(s-&gt;lexical != LNAME &amp;&amp; s-&gt;lexical != LTYPE){
			yyerror(&#34;invalid struct/union tag: %s&#34;, s-&gt;name);
			goto out;
		}
		dotag(s, et, 0);
		istag = 1;
	}else if(strcmp(s-&gt;name, &#34;_off_&#34;) == 0){
		debug[&#39;T&#39;] = 0;
		goto out;
	}else if(strcmp(s-&gt;name, &#34;_on_&#34;) == 0){
		debug[&#39;T&#39;] = 1;
		goto out;
	}
	t = s-&gt;type;
	if(istag)
		t = s-&gt;suetag;
	if(t == T)
		yyerror(&#34;unknown type %s in pragma incomplete&#34;, s-&gt;name);
	else if(!typesu[t-&gt;etype])
		yyerror(&#34;not struct/union type in pragma incomplete: %s&#34;, s-&gt;name);
	else
		t-&gt;garb |= GINCOMPLETE;
out:
	while(getnsc() != &#39;\n&#39;)
		;
	if(debug[&#39;f&#39;])
		print(&#34;%s incomplete\n&#34;, s-&gt;name);
}

void
pragdynld(void)
{
	Sym *local, *remote;
	char *path;
	Dynld *f;

	local = getsym();
	if(local == nil)
		goto err;

	remote = getsym();
	if(remote == nil)
		goto err;

	path = getquoted();
	if(path == nil)
		goto err;

	if(ndynld%32 == 0)
		dynld = realloc(dynld, (ndynld+32)*sizeof dynld[0]);
	f = &amp;dynld[ndynld++];
	f-&gt;local = local-&gt;name;
	f-&gt;remote = remote-&gt;name;
	f-&gt;path = path;
	goto out;

err:
	yyerror(&#34;usage: #pragma dynld local remote \&#34;path\&#34;&#34;);

out:
	while(getnsc() != &#39;\n&#39;)
		;
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
