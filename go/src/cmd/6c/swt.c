<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/6c/swt.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/6c/swt.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/6c/swt.c
// http://code.google.com/p/inferno-os/source/browse/utils/6c/swt.c
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

#include &#34;gc.h&#34;

void
swit1(C1 *q, int nc, int32 def, Node *n)
{
	C1 *r;
	int i;
	Prog *sp;

	if(nc &lt; 5) {
		for(i=0; i&lt;nc; i++) {
			if(debug[&#39;W&#39;])
				print(&#34;case = %.8llux\n&#34;, q-&gt;val);
			gcmp(OEQ, n, q-&gt;val);
			patch(p, q-&gt;label);
			q++;
		}
		gbranch(OGOTO);
		patch(p, def);
		return;
	}
	i = nc / 2;
	r = q+i;
	if(debug[&#39;W&#39;])
		print(&#34;case &gt; %.8llux\n&#34;, r-&gt;val);
	gcmp(OGT, n, r-&gt;val);
	sp = p;
	gbranch(OGOTO);
	p-&gt;as = AJEQ;
	patch(p, r-&gt;label);
	swit1(q, i, def, n);

	if(debug[&#39;W&#39;])
		print(&#34;case &lt; %.8llux\n&#34;, r-&gt;val);
	patch(sp, pc);
	swit1(r+1, nc-i-1, def, n);
}

void
bitload(Node *b, Node *n1, Node *n2, Node *n3, Node *nn)
{
	int sh;
	int32 v;
	Node *l;

	/*
	 * n1 gets adjusted/masked value
	 * n2 gets address of cell
	 * n3 gets contents of cell
	 */
	l = b-&gt;left;
	if(n2 != Z) {
		regalloc(n1, l, nn);
		reglcgen(n2, l, Z);
		regalloc(n3, l, Z);
		gmove(n2, n3);
		gmove(n3, n1);
	} else {
		regalloc(n1, l, nn);
		cgen(l, n1);
	}
	if(b-&gt;type-&gt;shift == 0 &amp;&amp; typeu[b-&gt;type-&gt;etype]) {
		v = ~0 + (1L &lt;&lt; b-&gt;type-&gt;nbits);
		gopcode(OAND, tfield, nodconst(v), n1);
	} else {
		sh = 32 - b-&gt;type-&gt;shift - b-&gt;type-&gt;nbits;
		if(sh &gt; 0)
			gopcode(OASHL, tfield, nodconst(sh), n1);
		sh += b-&gt;type-&gt;shift;
		if(sh &gt; 0)
			if(typeu[b-&gt;type-&gt;etype])
				gopcode(OLSHR, tfield, nodconst(sh), n1);
			else
				gopcode(OASHR, tfield, nodconst(sh), n1);
	}
}

void
bitstore(Node *b, Node *n1, Node *n2, Node *n3, Node *nn)
{
	int32 v;
	Node nod;
	int sh;

	regalloc(&amp;nod, b-&gt;left, Z);
	v = ~0 + (1L &lt;&lt; b-&gt;type-&gt;nbits);
	gopcode(OAND, types[TLONG], nodconst(v), n1);
	gmove(n1, &amp;nod);
	if(nn != Z)
		gmove(n1, nn);
	sh = b-&gt;type-&gt;shift;
	if(sh &gt; 0)
		gopcode(OASHL, types[TLONG], nodconst(sh), &amp;nod);
	v &lt;&lt;= sh;
	gopcode(OAND, types[TLONG], nodconst(~v), n3);
	gopcode(OOR, types[TLONG], n3, &amp;nod);
	gmove(&amp;nod, n2);

	regfree(&amp;nod);
	regfree(n1);
	regfree(n2);
	regfree(n3);
}

int32
outstring(char *s, int32 n)
{
	int32 r;

	if(suppress)
		return nstring;
	r = nstring;
	while(n) {
		string[mnstring] = *s++;
		mnstring++;
		nstring++;
		if(mnstring &gt;= NSNAME) {
			gpseudo(ADATA, symstring, nodconst(0L));
			p-&gt;from.offset += nstring - NSNAME;
			p-&gt;from.scale = NSNAME;
			p-&gt;to.type = D_SCONST;
			memmove(p-&gt;to.sval, string, NSNAME);
			mnstring = 0;
		}
		n--;
	}
	return r;
}

void
sextern(Sym *s, Node *a, int32 o, int32 w)
{
	int32 e, lw;

	for(e=0; e&lt;w; e+=NSNAME) {
		lw = NSNAME;
		if(w-e &lt; lw)
			lw = w-e;
		gpseudo(ADATA, s, nodconst(0L));
		p-&gt;from.offset += o+e;
		p-&gt;from.scale = lw;
		p-&gt;to.type = D_SCONST;
		memmove(p-&gt;to.sval, a-&gt;cstring+e, lw);
	}
}

void
gextern(Sym *s, Node *a, int32 o, int32 w)
{
	if(0 &amp;&amp; a-&gt;op == OCONST &amp;&amp; typev[a-&gt;type-&gt;etype]) {
		gpseudo(ADATA, s, lo64(a));
		p-&gt;from.offset += o;
		p-&gt;from.scale = 4;
		gpseudo(ADATA, s, hi64(a));
		p-&gt;from.offset += o + 4;
		p-&gt;from.scale = 4;
		return;
	}
	gpseudo(ADATA, s, a);
	p-&gt;from.offset += o;
	p-&gt;from.scale = w;
	switch(p-&gt;to.type) {
	default:
		p-&gt;to.index = p-&gt;to.type;
		p-&gt;to.type = D_ADDR;
	case D_CONST:
	case D_FCONST:
	case D_ADDR:
		break;
	}
}

void	zname(Biobuf*, Sym*, int);
void	zaddr(Biobuf*, Adr*, int);
void	outhist(Biobuf*);

void
outcode(void)
{
	struct { Sym *sym; short type; } h[NSYM];
	Prog *p;
	Sym *s;
	int f, sf, st, t, sym;
	Biobuf b;

	if(debug[&#39;S&#39;]) {
		for(p = firstp; p != P; p = p-&gt;link)
			if(p-&gt;as != ADATA &amp;&amp; p-&gt;as != AGLOBL)
				pc--;
		for(p = firstp; p != P; p = p-&gt;link) {
			print(&#34;%P\n&#34;, p);
			if(p-&gt;as != ADATA &amp;&amp; p-&gt;as != AGLOBL)
				pc++;
		}
	}

	f = open(outfile, OWRITE);
	if(f &lt; 0) {
		diag(Z, &#34;cannot open %s&#34;, outfile);
		return;
	}
	Binit(&amp;b, f, OWRITE);

	Bprint(&amp;b, &#34;%s\n&#34;, thestring);
	if(ndynld &gt; 0) {
		int i;

		Bprint(&amp;b, &#34;\n&#34;);
		Bprint(&amp;b, &#34;$$  // exports\n\n&#34;);
		Bprint(&amp;b, &#34;$$  // local types\n\n&#34;);
		Bprint(&amp;b, &#34;$$  // dynld\n&#34;, thestring);
		for(i=0; i&lt;ndynld; i++)
			Bprint(&amp;b, &#34;dynld %s %s %s\n&#34;, dynld[i].local, dynld[i].remote, dynld[i].path);
		Bprint(&amp;b, &#34;$$\n\n&#34;);
	}
	Bprint(&amp;b, &#34;!\n&#34;);

	outhist(&amp;b);
	for(sym=0; sym&lt;NSYM; sym++) {
		h[sym].sym = S;
		h[sym].type = 0;
	}
	sym = 1;
	for(p = firstp; p != P; p = p-&gt;link) {
	jackpot:
		sf = 0;
		s = p-&gt;from.sym;
		while(s != S) {
			sf = s-&gt;sym;
			if(sf &lt; 0 || sf &gt;= NSYM)
				sf = 0;
			t = p-&gt;from.type;
			if(t == D_ADDR)
				t = p-&gt;from.index;
			if(h[sf].type == t)
			if(h[sf].sym == s)
				break;
			s-&gt;sym = sym;
			zname(&amp;b, s, t);
			h[sym].sym = s;
			h[sym].type = t;
			sf = sym;
			sym++;
			if(sym &gt;= NSYM)
				sym = 1;
			break;
		}
		st = 0;
		s = p-&gt;to.sym;
		while(s != S) {
			st = s-&gt;sym;
			if(st &lt; 0 || st &gt;= NSYM)
				st = 0;
			t = p-&gt;to.type;
			if(t == D_ADDR)
				t = p-&gt;to.index;
			if(h[st].type == t)
			if(h[st].sym == s)
				break;
			s-&gt;sym = sym;
			zname(&amp;b, s, t);
			h[sym].sym = s;
			h[sym].type = t;
			st = sym;
			sym++;
			if(sym &gt;= NSYM)
				sym = 1;
			if(st == sf)
				goto jackpot;
			break;
		}
		Bputc(&amp;b, p-&gt;as);
		Bputc(&amp;b, p-&gt;as&gt;&gt;8);
		Bputc(&amp;b, p-&gt;lineno);
		Bputc(&amp;b, p-&gt;lineno&gt;&gt;8);
		Bputc(&amp;b, p-&gt;lineno&gt;&gt;16);
		Bputc(&amp;b, p-&gt;lineno&gt;&gt;24);
		zaddr(&amp;b, &amp;p-&gt;from, sf);
		zaddr(&amp;b, &amp;p-&gt;to, st);
	}
	Bflush(&amp;b);
	close(f);
	firstp = P;
	lastp = P;
}

void
outhist(Biobuf *b)
{
	Hist *h;
	char *p, *q, *op, c;
	Prog pg;
	int n;

	pg = zprog;
	pg.as = AHISTORY;
	c = pathchar();
	for(h = hist; h != H; h = h-&gt;link) {
		p = h-&gt;name;
		op = 0;
		/* on windows skip drive specifier in pathname */
		if(systemtype(Windows) &amp;&amp; p &amp;&amp; p[1] == &#39;:&#39;){
			p += 2;
			c = *p;
		}
		if(p &amp;&amp; p[0] != c &amp;&amp; h-&gt;offset == 0 &amp;&amp; pathname){
			/* on windows skip drive specifier in pathname */
			if(systemtype(Windows) &amp;&amp; pathname[1] == &#39;:&#39;) {
				op = p;
				p = pathname+2;
				c = *p;
			} else if(pathname[0] == c){
				op = p;
				p = pathname;
			}
		}
		while(p) {
			q = utfrune(p, c);
			if(q) {
				n = q-p;
				if(n == 0){
					n = 1;	/* leading &#34;/&#34; */
					*p = &#39;/&#39;;	/* don&#39;t emit &#34;\&#34; on windows */
				}
				q++;
			} else {
				n = strlen(p);
				q = 0;
			}
			if(n) {
				Bputc(b, ANAME);
				Bputc(b, ANAME&gt;&gt;8);
				Bputc(b, D_FILE);
				Bputc(b, 1);
				Bputc(b, &#39;&lt;&#39;);
				Bwrite(b, p, n);
				Bputc(b, 0);
			}
			p = q;
			if(p == 0 &amp;&amp; op) {
				p = op;
				op = 0;
			}
		}
		pg.lineno = h-&gt;line;
		pg.to.type = zprog.to.type;
		pg.to.offset = h-&gt;offset;
		if(h-&gt;offset)
			pg.to.type = D_CONST;

		Bputc(b, pg.as);
		Bputc(b, pg.as&gt;&gt;8);
		Bputc(b, pg.lineno);
		Bputc(b, pg.lineno&gt;&gt;8);
		Bputc(b, pg.lineno&gt;&gt;16);
		Bputc(b, pg.lineno&gt;&gt;24);
		zaddr(b, &amp;pg.from, 0);
		zaddr(b, &amp;pg.to, 0);
	}
}

void
zname(Biobuf *b, Sym *s, int t)
{
	char *n;
	uint32 sig;

	if(debug[&#39;T&#39;] &amp;&amp; t == D_EXTERN &amp;&amp; s-&gt;sig != SIGDONE &amp;&amp; s-&gt;type != types[TENUM] &amp;&amp; s != symrathole){
		sig = sign(s);
		Bputc(b, ASIGNAME);
		Bputc(b, ASIGNAME&gt;&gt;8);
		Bputc(b, sig);
		Bputc(b, sig&gt;&gt;8);
		Bputc(b, sig&gt;&gt;16);
		Bputc(b, sig&gt;&gt;24);
		s-&gt;sig = SIGDONE;
	}
	else{
		Bputc(b, ANAME);	/* as */
		Bputc(b, ANAME&gt;&gt;8);	/* as */
	}
	Bputc(b, t);			/* type */
	Bputc(b, s-&gt;sym);		/* sym */
	n = s-&gt;name;
	while(*n) {
		Bputc(b, *n);
		n++;
	}
	Bputc(b, 0);
}

void
zaddr(Biobuf *b, Adr *a, int s)
{
	int32 l;
	int i, t;
	char *n;
	Ieee e;

	t = 0;
	if(a-&gt;index != D_NONE || a-&gt;scale != 0)
		t |= T_INDEX;
	if(s != 0)
		t |= T_SYM;

	switch(a-&gt;type) {
	default:
		t |= T_TYPE;
	case D_NONE:
		if(a-&gt;offset != 0) {
			t |= T_OFFSET;
			l = a-&gt;offset;
			if((vlong)l != a-&gt;offset)
				t |= T_64;
		}
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
		if(t &amp; T_64) {
			l = a-&gt;offset&gt;&gt;32;
			Bputc(b, l);
			Bputc(b, l&gt;&gt;8);
			Bputc(b, l&gt;&gt;16);
			Bputc(b, l&gt;&gt;24);
		}
	}
	if(t &amp; T_SYM)		/* implies sym */
		Bputc(b, s);
	if(t &amp; T_FCONST) {
		ieeedtod(&amp;e, a-&gt;dval);
		l = e.l;
		Bputc(b, l);
		Bputc(b, l&gt;&gt;8);
		Bputc(b, l&gt;&gt;16);
		Bputc(b, l&gt;&gt;24);
		l = e.h;
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
}

int32
align(int32 i, Type *t, int op)
{
	int32 o;
	Type *v;
	int w;

	o = i;
	w = 1;
	switch(op) {
	default:
		diag(Z, &#34;unknown align opcode %d&#34;, op);
		break;

	case Asu2:	/* padding at end of a struct */
		w = SZ_VLONG;
		if(packflg)
			w = packflg;
		break;

	case Ael1:	/* initial align of struct element */
		for(v=t; v-&gt;etype==TARRAY; v=v-&gt;link)
			;
		w = ewidth[v-&gt;etype];
		if(w &lt;= 0 || w &gt;= SZ_VLONG)
			w = SZ_VLONG;
		if(packflg)
			w = packflg;
		break;

	case Ael2:	/* width of a struct element */
		o += t-&gt;width;
		break;

	case Aarg0:	/* initial passbyptr argument in arg list */
		if(typesu[t-&gt;etype]) {
			o = align(o, types[TIND], Aarg1);
			o = align(o, types[TIND], Aarg2);
		}
		break;

	case Aarg1:	/* initial align of parameter */
		w = ewidth[t-&gt;etype];
		if(w &lt;= 0 || w &gt;= SZ_VLONG) {
			w = SZ_VLONG;
			break;
		}
		w = 1;		/* little endian no adjustment */
		break;

	case Aarg2:	/* width of a parameter */
		o += t-&gt;width;
		w = t-&gt;width;
		if(w &gt; SZ_VLONG)
			w = SZ_VLONG;
		break;

	case Aaut3:	/* total allign of automatic */
		o = align(o, t, Ael1);
		o = align(o, t, Ael2);
		break;
	}
	o = xround(o, w);
	if(debug[&#39;A&#39;])
		print(&#34;align %s %ld %T = %ld\n&#34;, bnames[op], i, t, o);
	return o;
}

int32
maxround(int32 max, int32 v)
{
	v += SZ_VLONG-1;
	if(v &gt; max)
		max = xround(v, SZ_VLONG);
	return max;
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
