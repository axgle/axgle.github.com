<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8g/gobj.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/8g/gobj.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Derived from Inferno utils/8c/swt.c
// http://code.google.com/p/inferno-os/source/browse/utils/8c/swt.c
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

#include &#34;gg.h&#34;

void
zname(Biobuf *b, Sym *s, int t)
{
	char *n;

	Bputc(b, ANAME);	/* as */
	Bputc(b, ANAME&gt;&gt;8);	/* as */
	Bputc(b, t);		/* type */
	Bputc(b, s-&gt;sym);	/* sym */

	for(n=s-&gt;package; *n; n++)
		Bputc(b, *n);
	Bputdot(b);
	for(n=s-&gt;name; *n; n++)
		Bputc(b, *n);
	Bputc(b, 0);
}

void
zfile(Biobuf *b, char *p, int n)
{
	Bputc(b, ANAME);
	Bputc(b, ANAME&gt;&gt;8);
	Bputc(b, D_FILE);
	Bputc(b, 1);
	Bputc(b, &#39;&lt;&#39;);
	Bwrite(b, p, n);
	Bputc(b, 0);
}

void
zhist(Biobuf *b, int line, vlong offset)
{
	Addr a;

	Bputc(b, AHISTORY);
	Bputc(b, AHISTORY&gt;&gt;8);
	Bputc(b, line);
	Bputc(b, line&gt;&gt;8);
	Bputc(b, line&gt;&gt;16);
	Bputc(b, line&gt;&gt;24);
	zaddr(b, &amp;zprog.from, 0, 0);
	a = zprog.to;
	if(offset != 0) {
		a.offset = offset;
		a.type = D_CONST;
	}
	zaddr(b, &amp;a, 0, 0);
}

void
zaddr(Biobuf *b, Addr *a, int s, int gotype)
{
	int32 l;
	uint64 e;
	int i, t;
	char *n;

	t = 0;
	if(a-&gt;index != D_NONE || a-&gt;scale != 0)
		t |= T_INDEX;
	if(s != 0)
		t |= T_SYM;
	if(gotype != 0)
		t |= T_GOTYPE;

	switch(a-&gt;type) {

	case D_BRANCH:
		if(a-&gt;branch == nil)
			fatal(&#34;unpatched branch&#34;);
		a-&gt;offset = a-&gt;branch-&gt;loc;

	default:
		t |= T_TYPE;

	case D_NONE:
		if(a-&gt;offset != 0)
			t |= T_OFFSET;
		if(a-&gt;offset2 != 0)
			t |= T_OFFSET2;
		break;
	case D_FCONST:
		t |= T_FCONST;
		break;
	case D_SCONST:
		t |= T_SCONST;
		break;
	}
	Bputc(b, t);

	if(t &amp; T_INDEX) {	/* implies index, scale */
		Bputc(b, a-&gt;index);
		Bputc(b, a-&gt;scale);
	}
	if(t &amp; T_OFFSET) {	/* implies offset */
		l = a-&gt;offset;
		Bputc(b, l);
		Bputc(b, l&gt;&gt;8);
		Bputc(b, l&gt;&gt;16);
		Bputc(b, l&gt;&gt;24);
	}
	if(t &amp; T_OFFSET2) {	/* implies offset */
		l = a-&gt;offset2;
		Bputc(b, l);
		Bputc(b, l&gt;&gt;8);
		Bputc(b, l&gt;&gt;16);
		Bputc(b, l&gt;&gt;24);
	}
	if(t &amp; T_SYM)		/* implies sym */
		Bputc(b, s);
	if(t &amp; T_FCONST) {
		ieeedtod(&amp;e, a-&gt;dval);
		l = e;
		Bputc(b, l);
		Bputc(b, l&gt;&gt;8);
		Bputc(b, l&gt;&gt;16);
		Bputc(b, l&gt;&gt;24);
		l = e &gt;&gt; 32;
		Bputc(b, l);
		Bputc(b, l&gt;&gt;8);
		Bputc(b, l&gt;&gt;16);
		Bputc(b, l&gt;&gt;24);
		return;
	}
	if(t &amp; T_SCONST) {
		n = a-&gt;sval;
		for(i=0; i&lt;NSNAME; i++) {
			Bputc(b, *n);
			n++;
		}
		return;
	}
	if(t &amp; T_TYPE)
		Bputc(b, a-&gt;type);
	if(t &amp; T_GOTYPE)
		Bputc(b, gotype);
}

static struct {
	struct { Sym *sym; short type; } h[NSYM];
	int sym;
} z;

static void
zsymreset(void)
{
	for(z.sym=0; z.sym&lt;NSYM; z.sym++) {
		z.h[z.sym].sym = S;
		z.h[z.sym].type = 0;
	}
	z.sym = 1;
}

static int
zsym(Sym *s, int t, int *new)
{
	int i;

	*new = 0;
	if(s == S)
		return 0;

	i = s-&gt;sym;
	if(i &lt; 0 || i &gt;= NSYM)
		i = 0;
	if(z.h[i].type == t &amp;&amp; z.h[i].sym == s)
		return i;
	i = z.sym;
	s-&gt;sym = i;
	zname(bout, s, t);
	z.h[i].sym = s;
	z.h[i].type = t;
	if(++z.sym &gt;= NSYM)
		z.sym = 1;
	*new = 1;
	return i;
}

static int
zsymaddr(Addr *a, int *new)
{
	int t;

	t = a-&gt;type;
	if(t == D_ADDR)
		t = a-&gt;index;
	return zsym(a-&gt;sym, t, new);
}

void
dumpfuncs(void)
{
	Plist *pl;
	int sf, st, gf, gt, new;
	Sym *s;
	Prog *p;

	zsymreset();

	// fix up pc
	pcloc = 0;
	for(pl=plist; pl!=nil; pl=pl-&gt;link) {
		for(p=pl-&gt;firstpc; p!=P; p=p-&gt;link) {
			p-&gt;loc = pcloc;
			if(p-&gt;as != ADATA &amp;&amp; p-&gt;as != AGLOBL)
				pcloc++;
		}
	}

	// put out functions
	for(pl=plist; pl!=nil; pl=pl-&gt;link) {

		if(debug[&#39;S&#39;]) {
			s = S;
			if(pl-&gt;name != N)
				s = pl-&gt;name-&gt;sym;
			print(&#34;\n--- prog list \&#34;%S\&#34; ---\n&#34;, s);
			for(p=pl-&gt;firstpc; p!=P; p=p-&gt;link)
				print(&#34;%P\n&#34;, p);
		}

		for(p=pl-&gt;firstpc; p!=P; p=p-&gt;link) {
			for(;;) {
				sf = zsymaddr(&amp;p-&gt;from, &amp;new);
				gf = zsym(p-&gt;from.gotype, D_EXTERN, &amp;new);
				if(new &amp;&amp; sf == gf)
					continue;
				st = zsymaddr(&amp;p-&gt;to, &amp;new);
				if(new &amp;&amp; (st == sf || st == gf))
					continue;
				gt = zsym(p-&gt;to.gotype, D_EXTERN, &amp;new);
				if(new &amp;&amp; (gt == sf || gt == gf || gt == st))
					continue;
				break;
			}

			Bputc(bout, p-&gt;as);
			Bputc(bout, p-&gt;as&gt;&gt;8);
			Bputc(bout, p-&gt;lineno);
			Bputc(bout, p-&gt;lineno&gt;&gt;8);
			Bputc(bout, p-&gt;lineno&gt;&gt;16);
			Bputc(bout, p-&gt;lineno&gt;&gt;24);
			zaddr(bout, &amp;p-&gt;from, sf, gf);
			zaddr(bout, &amp;p-&gt;to, st, gt);
		}
	}
}

/* deferred DATA output */
static Prog *strdat;
static Prog *estrdat;
static int gflag;
static Prog *savepc;

void
data(void)
{
	gflag = debug[&#39;g&#39;];
	debug[&#39;g&#39;] = 0;

	if(estrdat == nil) {
		strdat = mal(sizeof(*pc));
		clearp(strdat);
		estrdat = strdat;
	}
	if(savepc)
		fatal(&#34;data phase error&#34;);
	savepc = pc;
	pc = estrdat;
}

void
text(void)
{
	if(!savepc)
		fatal(&#34;text phase error&#34;);
	debug[&#39;g&#39;] = gflag;
	estrdat = pc;
	pc = savepc;
	savepc = nil;
}

void
dumpdata(void)
{
	Prog *p;

	if(estrdat == nil)
		return;
	*pc = *strdat;
	if(gflag)
		for(p=pc; p!=estrdat; p=p-&gt;link)
			print(&#34;%P\n&#34;, p);
	pc = estrdat;
}

/*
 * make a refer to the data s, s+len
 * emitting DATA if needed.
 */
void
datastring(char *s, int len, Addr *a)
{
	int w;
	Prog *p;
	Addr ac, ao;
	static int gen;
	struct {
		Strlit lit;
		char buf[100];
	} tmp;

	// string
	memset(&amp;ao, 0, sizeof(ao));
	ao.type = D_STATIC;
	ao.index = D_NONE;
	ao.etype = TINT32;
	ao.offset = 0;		// fill in

	// constant
	memset(&amp;ac, 0, sizeof(ac));
	ac.type = D_CONST;
	ac.index = D_NONE;
	ac.offset = 0;		// fill in

	// huge strings are made static to avoid long names.
	if(len &gt; 100) {
		snprint(namebuf, sizeof(namebuf), &#34;.string.%d&#34;, gen++);
		ao.sym = lookup(namebuf);
		ao.type = D_STATIC;
	} else {
		if(len &gt; 0 &amp;&amp; s[len-1] == &#39;\0&#39;)
			len--;
		tmp.lit.len = len;
		memmove(tmp.lit.s, s, len);
		tmp.lit.s[len] = &#39;\0&#39;;
		len++;
		snprint(namebuf, sizeof(namebuf), &#34;\&#34;%Z\&#34;&#34;, &amp;tmp.lit);
		ao.sym = pkglookup(namebuf, &#34;string&#34;);
		ao.type = D_EXTERN;
	}
	*a = ao;

	// only generate data the first time.
	if(ao.sym-&gt;flags &amp; SymUniq)
		return;
	ao.sym-&gt;flags |= SymUniq;

	data();
	for(w=0; w&lt;len; w+=8) {
		p = pc;
		gins(ADATA, N, N);

		// DATA s+w, [NSNAME], $&#34;xxx&#34;
		p-&gt;from = ao;
		p-&gt;from.offset = w;

		p-&gt;from.scale = NSNAME;
		if(w+8 &gt; len)
			p-&gt;from.scale = len-w;

		p-&gt;to = ac;
		p-&gt;to.type = D_SCONST;
		p-&gt;to.offset = len;
		memmove(p-&gt;to.sval, s+w, p-&gt;from.scale);
	}
	p = pc;
	ggloblsym(ao.sym, len, ao.type == D_EXTERN);
	if(ao.type == D_STATIC)
		p-&gt;from.type = D_STATIC;
	text();
}

/*
 * make a refer to the string sval,
 * emitting DATA if needed.
 */
void
datagostring(Strlit *sval, Addr *a)
{
	Prog *p;
	Addr ac, ao, ap;
	int32 wi, wp;
	static int gen;

	memset(&amp;ac, 0, sizeof(ac));
	memset(&amp;ao, 0, sizeof(ao));
	memset(&amp;ap, 0, sizeof(ap));

	// constant
	ac.type = D_CONST;
	ac.index = D_NONE;
	ac.offset = 0;			// fill in

	// string len+ptr
	ao.type = D_STATIC;		// fill in
	ao.index = D_NONE;
	ao.etype = TINT32;
	ao.sym = nil;			// fill in

	// $string len+ptr
	datastring(sval-&gt;s, sval-&gt;len, &amp;ap);
	ap.index = ap.type;
	ap.type = D_ADDR;
	ap.etype = TINT32;

	wi = types[TUINT32]-&gt;width;
	wp = types[tptr]-&gt;width;

	if(ap.index == D_STATIC) {
		// huge strings are made static to avoid long names
		snprint(namebuf, sizeof(namebuf), &#34;.gostring.%d&#34;, ++gen);
		ao.sym = lookup(namebuf);
		ao.type = D_STATIC;
	} else {
		// small strings get named by their contents,
		// so that multiple modules using the same string
		// can share it.
		snprint(namebuf, sizeof(namebuf), &#34;\&#34;%Z\&#34;&#34;, sval);
		ao.sym = pkglookup(namebuf, &#34;go.string&#34;);
		ao.type = D_EXTERN;
	}

	*a = ao;
	if(ao.sym-&gt;flags &amp; SymUniq)
		return;
	ao.sym-&gt;flags |= SymUniq;

	data();
	// DATA gostring, wp, $cstring
	p = pc;
	gins(ADATA, N, N);
	p-&gt;from = ao;
	p-&gt;from.scale = wp;
	p-&gt;to = ap;

	// DATA gostring+wp, wi, $len
	p = pc;
	gins(ADATA, N, N);
	p-&gt;from = ao;
	p-&gt;from.offset = wp;
	p-&gt;from.scale = wi;
	p-&gt;to = ac;
	p-&gt;to.offset = sval-&gt;len;

	p = pc;
	ggloblsym(ao.sym, types[TSTRING]-&gt;width, ao.type == D_EXTERN);
	if(ao.type == D_STATIC)
		p-&gt;from.type = D_STATIC;
	text();
}

void
gdata(Node *nam, Node *nr, int wid)
{
	Prog *p;
	vlong v;

	if(wid == 8 &amp;&amp; is64(nr-&gt;type)) {
		v = mpgetfix(nr-&gt;val.u.xval);
		p = gins(ADATA, nam, nodintconst(v));
		p-&gt;from.scale = 4;
		p = gins(ADATA, nam, nodintconst(v&gt;&gt;32));
		p-&gt;from.scale = 4;
		p-&gt;from.offset += 4;
		return;
	}
	p = gins(ADATA, nam, nr);
	p-&gt;from.scale = wid;
}

void
gdatastring(Node *nam, Strlit *sval)
{
	Prog *p;
	Node nod1;

	p = gins(ADATA, nam, N);
	datastring(sval-&gt;s, sval-&gt;len, &amp;p-&gt;to);
	p-&gt;from.scale = types[tptr]-&gt;width;
	p-&gt;to.index = p-&gt;to.type;
	p-&gt;to.type = D_ADDR;
//print(&#34;%P\n&#34;, p);

	nodconst(&amp;nod1, types[TINT32], sval-&gt;len);
	p = gins(ADATA, nam, &amp;nod1);
	p-&gt;from.scale = types[TINT32]-&gt;width;
	p-&gt;from.offset += types[tptr]-&gt;width;
}

int
dstringptr(Sym *s, int off, char *str)
{
	Prog *p;

	off = rnd(off, widthptr);
	p = gins(ADATA, N, N);
	p-&gt;from.type = D_EXTERN;
	p-&gt;from.index = D_NONE;
	p-&gt;from.sym = s;
	p-&gt;from.offset = off;
	p-&gt;from.scale = widthptr;

	datastring(str, strlen(str)+1, &amp;p-&gt;to);
	p-&gt;to.index = p-&gt;to.type;
	p-&gt;to.type = D_ADDR;
	p-&gt;to.etype = TINT32;
	off += widthptr;

	return off;
}

int
dgostrlitptr(Sym *s, int off, Strlit *lit)
{
	Prog *p;

	if(lit == nil)
		return duintptr(s, off, 0);

	off = rnd(off, widthptr);
	p = gins(ADATA, N, N);
	p-&gt;from.type = D_EXTERN;
	p-&gt;from.index = D_NONE;
	p-&gt;from.sym = s;
	p-&gt;from.offset = off;
	p-&gt;from.scale = widthptr;
	datagostring(lit, &amp;p-&gt;to);
	p-&gt;to.index = p-&gt;to.type;
	p-&gt;to.type = D_ADDR;
	p-&gt;to.etype = TINT32;
	off += widthptr;

	return off;
}

int
dgostringptr(Sym *s, int off, char *str)
{
	int n;
	Strlit *lit;

	if(str == nil)
		return duintptr(s, off, 0);

	n = strlen(str);
	lit = mal(sizeof *lit + n);
	strcpy(lit-&gt;s, str);
	lit-&gt;len = n;
	return dgostrlitptr(s, off, lit);
}


int
duintxx(Sym *s, int off, uint64 v, int wid)
{
	Prog *p;

	off = rnd(off, wid);

	p = gins(ADATA, N, N);
	p-&gt;from.type = D_EXTERN;
	p-&gt;from.index = D_NONE;
	p-&gt;from.sym = s;
	p-&gt;from.offset = off;
	p-&gt;from.scale = wid;
	p-&gt;to.type = D_CONST;
	p-&gt;to.index = D_NONE;
	p-&gt;to.offset = v;
	off += wid;

	return off;
}

int
dsymptr(Sym *s, int off, Sym *x, int xoff)
{
	Prog *p;

	off = rnd(off, widthptr);

	p = gins(ADATA, N, N);
	p-&gt;from.type = D_EXTERN;
	p-&gt;from.index = D_NONE;
	p-&gt;from.sym = s;
	p-&gt;from.offset = off;
	p-&gt;from.scale = widthptr;
	p-&gt;to.type = D_ADDR;
	p-&gt;to.index = D_EXTERN;
	p-&gt;to.sym = x;
	p-&gt;to.offset = xoff;
	off += widthptr;

	return off;
}

void
genembedtramp(Type *rcvr, Type *method, Sym *newnam)
{
	Sym *e;
	int c, d, o, mov, add, loaded;
	Prog *p;
	Type *f;

	e = method-&gt;sym;
	for(d=0; d&lt;nelem(dotlist); d++) {
		c = adddot1(e, rcvr, d, nil);
		if(c == 1)
			goto out;
	}
	fatal(&#34;genembedtramp %T.%S&#34;, rcvr, method-&gt;sym);

out:
	newplist()-&gt;name = newname(newnam);

	//TEXT	main·S_test2(SB),7,$0
	p = pc;
	gins(ATEXT, N, N);
	p-&gt;from.type = D_EXTERN;
	p-&gt;from.sym = newnam;
	p-&gt;to.type = D_CONST;
	p-&gt;to.offset = 0;
	p-&gt;from.scale = 7;
//print(&#34;1. %P\n&#34;, p);

	mov = AMOVL;
	add = AADDL;

	loaded = 0;
	o = 0;
	for(c=d-1; c&gt;=0; c--) {
		f = dotlist[c].field;
		o += f-&gt;width;
		if(!isptr[f-&gt;type-&gt;etype])
			continue;
		if(!loaded) {
			loaded = 1;
			//MOVL	4(SP), AX
			p = pc;
			gins(mov, N, N);
			p-&gt;from.type = D_INDIR+D_SP;
			p-&gt;from.offset = widthptr;
			p-&gt;to.type = D_AX;
//print(&#34;2. %P\n&#34;, p);
		}

		//MOVL	o(AX), AX
		p = pc;
		gins(mov, N, N);
		p-&gt;from.type = D_INDIR+D_AX;
		p-&gt;from.offset = o;
		p-&gt;to.type = D_AX;
//print(&#34;3. %P\n&#34;, p);
		o = 0;
	}
	if(o != 0) {
		//ADDL	$XX, AX
		p = pc;
		gins(add, N, N);
		p-&gt;from.type = D_CONST;
		p-&gt;from.offset = o;
		if(loaded)
			p-&gt;to.type = D_AX;
		else {
			p-&gt;to.type = D_INDIR+D_SP;
			p-&gt;to.offset = widthptr;
		}
//print(&#34;4. %P\n&#34;, p);
	}

	//MOVL	AX, 4(SP)
	if(loaded) {
		p = pc;
		gins(mov, N, N);
		p-&gt;from.type = D_AX;
		p-&gt;to.type = D_INDIR+D_SP;
		p-&gt;to.offset = widthptr;
//print(&#34;5. %P\n&#34;, p);
	} else {
		// TODO(rsc): obviously this is unnecessary,
		// but 6l has a bug, and it can&#39;t handle
		// JMP instructions too close to the top of
		// a new function.
		p = pc;
		gins(ANOP, N, N);
	}

	f = dotlist[0].field;
	//JMP	main·*Sub_test2(SB)
	if(isptr[f-&gt;type-&gt;etype])
		f = f-&gt;type;
	p = pc;
	gins(AJMP, N, N);
	p-&gt;to.type = D_EXTERN;
	p-&gt;to.sym = methodsym(method-&gt;sym, ptrto(f-&gt;type));
//print(&#34;6. %P\n&#34;, p);

	pc-&gt;as = ARET;	// overwrite AEND
}

void
nopout(Prog *p)
{
	p-&gt;as = ANOP;
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
