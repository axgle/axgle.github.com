<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5g/gsubr.c</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5g/gsubr.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Derived from Inferno utils/5c/txt.c
// http://code.google.com/p/inferno-os/source/browse/utils/5c/txt.c
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

// TODO(kaib): Can make this bigger if we move
// the text segment up higher in 5l for all GOOS.
long unmappedzero = 4096;

void
clearp(Prog *p)
{
	p-&gt;as = AEND;
	p-&gt;reg = NREG;
	p-&gt;scond = C_SCOND_NONE;
	p-&gt;from.type = D_NONE;
	p-&gt;from.name = D_NONE;
	p-&gt;from.reg = NREG;
	p-&gt;to.type = D_NONE;
	p-&gt;to.name = D_NONE;
	p-&gt;to.reg = NREG;
	p-&gt;loc = pcloc;
	pcloc++;
}

/*
 * generate and return proc with p-&gt;as = as,
 * linked into program.  pc is next instruction.
 */
Prog*
prog(int as)
{
	Prog *p;

	p = pc;
	pc = mal(sizeof(*pc));

	clearp(pc);

	if(lineno == 0) {
		if(debug[&#39;K&#39;])
			warn(&#34;prog: line 0&#34;);
	}

	p-&gt;as = as;
	p-&gt;lineno = lineno;
	p-&gt;link = pc;
	return p;
}

/*
 * generate a branch.
 * t is ignored.
 */
Prog*
gbranch(int as, Type *t)
{
	Prog *p;

	p = prog(as);
	p-&gt;to.type = D_BRANCH;
	p-&gt;to.branch = P;
	return p;
}

/*
 * patch previous branch to jump to to.
 */
void
patch(Prog *p, Prog *to)
{
	if(p-&gt;to.type != D_BRANCH)
		fatal(&#34;patch: not a branch&#34;);
	p-&gt;to.branch = to;
	p-&gt;to.offset = to-&gt;loc;
}

/*
 * start a new Prog list.
 */
Plist*
newplist(void)
{
	Plist *pl;

	pl = mal(sizeof(*pl));
	if(plist == nil)
		plist = pl;
	else
		plast-&gt;link = pl;
	plast = pl;

	pc = mal(sizeof(*pc));
	clearp(pc);
	pl-&gt;firstpc = pc;

	return pl;
}

void
gused(Node *n)
{
	gins(ANOP, n, N);	// used
}

Prog*
gjmp(Prog *to)
{
	Prog *p;

	p = gbranch(AB, T);
	if(to != P)
		patch(p, to);
	return p;
}

void
ggloblnod(Node *nam, int32 width)
{
	Prog *p;

	p = gins(AGLOBL, nam, N);
	p-&gt;lineno = nam-&gt;lineno;
	p-&gt;to.sym = S;
	p-&gt;to.type = D_CONST;
	p-&gt;to.offset = width;
}

void
ggloblsym(Sym *s, int32 width, int dupok)
{
	Prog *p;

	p = gins(AGLOBL, N, N);
	p-&gt;from.type = D_OREG;
	p-&gt;from.name = D_EXTERN;
	p-&gt;from.sym = s;
	p-&gt;to.type = D_CONST;
	p-&gt;to.name = D_NONE;
	p-&gt;to.offset = width;
	if(dupok)
		p-&gt;reg = DUPOK;
}

int
isfat(Type *t)
{
	if(t != T)
	switch(t-&gt;etype) {
	case TSTRUCT:
	case TARRAY:
	case TSTRING:
	case TINTER:	// maybe remove later
	case TDDD:	// maybe remove later
		return 1;
	}
	return 0;
}

/*
 * naddr of func generates code for address of func.
 * if using opcode that can take address implicitly,
 * call afunclit to fix up the argument.
 * also fix up direct register references to be D_OREG.
 */
void
afunclit(Addr *a)
{
	if(a-&gt;type == D_CONST &amp;&amp; a-&gt;name == D_EXTERN || a-&gt;type == D_REG) {
		a-&gt;type = D_OREG;
	}
}

/*
 * allocate register of type t, leave in n.
 * if o != N, o is desired fixed register.
 * caller must regfree(n).
 */
void
regalloc(Node *n, Type *t, Node *o)
{
	int i, et, fixfree, floatfree;

	if(debug[&#39;r&#39;]) {
		fixfree = 0;
		for(i=REGALLOC_R0; i&lt;=REGALLOC_RMAX; i++)
			if(reg[i] == 0)
				fixfree++;
		floatfree = 0;
		for(i=REGALLOC_F0; i&lt;=REGALLOC_FMAX; i++)
			if(reg[i] == 0)
				floatfree++;
		print(&#34;regalloc fix %d float %d\n&#34;, fixfree, floatfree);
	}

	if(t == T)
		fatal(&#34;regalloc: t nil&#34;);
	et = simtype[t-&gt;etype];
	if(is64(t))
		fatal(&#34;regalloc: 64 bit type %T&#34;);

	switch(et) {
	case TINT8:
	case TUINT8:
	case TINT16:
	case TUINT16:
	case TINT32:
	case TUINT32:
	case TPTR32:
	case TBOOL:
		if(o != N &amp;&amp; o-&gt;op == OREGISTER) {
			i = o-&gt;val.u.reg;
			if(i &gt;= REGALLOC_R0 &amp;&amp; i &lt;= REGALLOC_RMAX)
				goto out;
		}
		for(i=REGALLOC_R0; i&lt;=REGALLOC_RMAX; i++)
			if(reg[i] == 0)
				goto out;

		yyerror(&#34;out of fixed registers&#34;);
		goto err;

	case TFLOAT32:
	case TFLOAT64:
		if(o != N &amp;&amp; o-&gt;op == OREGISTER) {
			i = o-&gt;val.u.reg;
			if(i &gt;= REGALLOC_F0 &amp;&amp; i &lt;= REGALLOC_FMAX)
				goto out;
		}
		for(i=REGALLOC_F0; i&lt;=REGALLOC_FMAX; i++)
			if(reg[i] == 0)
				goto out;
		yyerror(&#34;out of floating point registers&#34;);
		goto err;
	}
	yyerror(&#34;regalloc: unknown type %T&#34;, t);

err:
	nodreg(n, t, 0);
	return;

out:
	reg[i]++;
	nodreg(n, t, i);
}

void
regfree(Node *n)
{
	int i, fixfree, floatfree;

	if(debug[&#39;r&#39;]) {
		fixfree = 0;
		for(i=REGALLOC_R0; i&lt;=REGALLOC_RMAX; i++)
			if(reg[i] == 0)
				fixfree++;
		floatfree = 0;
		for(i=REGALLOC_F0; i&lt;=REGALLOC_FMAX; i++)
			if(reg[i] == 0)
				floatfree++;
		print(&#34;regalloc fix %d float %d\n&#34;, fixfree, floatfree);
	}

	if(n-&gt;op != OREGISTER &amp;&amp; n-&gt;op != OINDREG)
		fatal(&#34;regfree: not a register&#34;);
	i = n-&gt;val.u.reg;
	if(i &lt; 0 || i &gt;= sizeof(reg))
		fatal(&#34;regfree: reg out of range&#34;);
	if(reg[i] &lt;= 0)
		fatal(&#34;regfree: reg not allocated&#34;);
	reg[i]--;
}

/*
 * initialize n to be register r of type t.
 */
void
nodreg(Node *n, Type *t, int r)
{
	if(t == T)
		fatal(&#34;nodreg: t nil&#34;);

	memset(n, 0, sizeof(*n));
	n-&gt;op = OREGISTER;
	n-&gt;addable = 1;
	ullmancalc(n);
	n-&gt;val.u.reg = r;
	n-&gt;type = t;
}

/*
 * initialize n to be indirect of register r; n is type t.
 */
void
nodindreg(Node *n, Type *t, int r)
{
	nodreg(n, t, r);
	n-&gt;op = OINDREG;
}

Node*
nodarg(Type *t, int fp)
{
	Node *n;
	Type *first;
	Iter savet;

	// entire argument struct, not just one arg
	if(t-&gt;etype == TSTRUCT &amp;&amp; t-&gt;funarg) {
		n = nod(ONAME, N, N);
		n-&gt;sym = lookup(&#34;.args&#34;);
		n-&gt;type = t;
		first = structfirst(&amp;savet, &amp;t);
		if(first == nil)
			fatal(&#34;nodarg: bad struct&#34;);
		if(first-&gt;width == BADWIDTH)
			fatal(&#34;nodarg: offset not computed for %T&#34;, t);
		n-&gt;xoffset = first-&gt;width;
		n-&gt;addable = 1;
		goto fp;
	}

	if(t-&gt;etype != TFIELD)
		fatal(&#34;nodarg: not field %T&#34;, t);

	n = nod(ONAME, N, N);
	n-&gt;type = t-&gt;type;
	n-&gt;sym = t-&gt;sym;
	if(t-&gt;width == BADWIDTH)
		fatal(&#34;nodarg: offset not computed for %T&#34;, t);
	n-&gt;xoffset = t-&gt;width;
	n-&gt;addable = 1;

fp:
	switch(fp) {
	default:
		fatal(&#34;nodarg %T %d&#34;, t, fp);

	case 0:		// output arg for calling another function
		n-&gt;op = OINDREG;
		n-&gt;val.u.reg = REGSP;
		n-&gt;xoffset += 4;
		break;

	case 1:		// input arg to current function
		n-&gt;class = PPARAM;
		break;
	}
	return n;
}

/*
 * return constant i node.
 * overwritten by next call, but useful in calls to gins.
 */
Node*
ncon(uint32 i)
{
	static Node n;

	if(n.type == T)
		nodconst(&amp;n, types[TUINT32], 0);
	mpmovecfix(n.val.u.xval, i);
	return &amp;n;
}

/*
 * Is this node a memory operand?
 */
int
ismem(Node *n)
{
	switch(n-&gt;op) {
	case OINDREG:
	case ONAME:
	case OPARAM:
		return 1;
	}
	return 0;
}

Node sclean[10];
int nsclean;

/*
 * n is a 64-bit value.  fill in lo and hi to refer to its 32-bit halves.
 */
void
split64(Node *n, Node *lo, Node *hi)
{
	Node n1;
	int64 i;

	if(!is64(n-&gt;type))
		fatal(&#34;split64 %T&#34;, n-&gt;type);

	sclean[nsclean].op = OEMPTY;
	if(nsclean &gt;= nelem(sclean))
		fatal(&#34;split64 clean&#34;);
	nsclean++;
	switch(n-&gt;op) {
	default:
		if(!dotaddable(n, &amp;n1)) {
			igen(n, &amp;n1, N);
			sclean[nsclean-1] = n1;
		}
		n = &amp;n1;
		goto common;
	case ONAME:
		if(n-&gt;class == PPARAMREF) {
			cgen(n-&gt;heapaddr, &amp;n1);
			sclean[nsclean-1] = n1;
			// fall through.
			n = &amp;n1;
		}
		goto common;
	case OINDREG:
	common:
		*lo = *n;
		*hi = *n;
		lo-&gt;type = types[TUINT32];
		if(n-&gt;type-&gt;etype == TINT64)
			hi-&gt;type = types[TINT32];
		else
			hi-&gt;type = types[TUINT32];
		hi-&gt;xoffset += 4;
		break;

	case OLITERAL:
		convconst(&amp;n1, n-&gt;type, &amp;n-&gt;val);
		i = mpgetfix(n1.val.u.xval);
		nodconst(lo, types[TUINT32], (uint32)i);
		i &gt;&gt;= 32;
		if(n-&gt;type-&gt;etype == TINT64)
			nodconst(hi, types[TINT32], (int32)i);
		else
			nodconst(hi, types[TUINT32], (uint32)i);
		break;
	}
}

void
splitclean(void)
{
	if(nsclean &lt;= 0)
		fatal(&#34;splitclean&#34;);
	nsclean--;
	if(sclean[nsclean].op != OEMPTY)
		regfree(&amp;sclean[nsclean]);
}

#define	CASE(a,b)	(((a)&lt;&lt;16)|((b)&lt;&lt;0))

void
gmove(Node *f, Node *t)
{
	int a, ft, tt, fa, ta;
	Type *cvt;
	Node r1, r2, flo, fhi, tlo, thi, con;
	Prog *p1;

	if(debug[&#39;M&#39;])
		print(&#34;gmove %N -&gt; %N\n&#34;, f, t);

	ft = simsimtype(f-&gt;type);
	tt = simsimtype(t-&gt;type);
	cvt = t-&gt;type;

	// cannot have two memory operands;
	// except 64-bit, which always copies via registers anyway.
	if(!is64(f-&gt;type) &amp;&amp; !is64(t-&gt;type) &amp;&amp; ismem(f) &amp;&amp; ismem(t))
		goto hard;

	// convert constant to desired type
	if(f-&gt;op == OLITERAL) {
		switch(tt) {
		default:
			convconst(&amp;con, t-&gt;type, &amp;f-&gt;val);
			break;

		case TINT16:
		case TINT8:
			convconst(&amp;con, types[TINT32], &amp;f-&gt;val);
			regalloc(&amp;r1, con.type, t);
			gins(AMOVW, &amp;con, &amp;r1);
			gmove(&amp;r1, t);
			regfree(&amp;r1);
			return;

		case TUINT16:
		case TUINT8:
			convconst(&amp;con, types[TUINT32], &amp;f-&gt;val);
			regalloc(&amp;r1, con.type, t);
			gins(AMOVW, &amp;con, &amp;r1);
			gmove(&amp;r1, t);
			regfree(&amp;r1);
			return;
		}

		f = &amp;con;
		ft = simsimtype(con.type);

		// constants can&#39;t move directly to memory
		if(ismem(t) &amp;&amp; !is64(t-&gt;type)) goto hard;
	}

	// value -&gt; value copy, only one memory operand.
	// figure out the instruction to use.
	// break out of switch for one-instruction gins.
	// goto rdst for &#34;destination must be register&#34;.
	// goto hard for &#34;convert to cvt type first&#34;.
	// otherwise handle and return.

	switch(CASE(ft, tt)) {
	default:
		goto fatal;

	/*
	 * integer copy and truncate
	 */
	case CASE(TINT8, TINT8):	// same size
	case CASE(TUINT8, TINT8):
	case CASE(TINT16, TINT8):	// truncate
	case CASE(TUINT16, TINT8):
	case CASE(TINT32, TINT8):
	case CASE(TUINT32, TINT8):
		a = AMOVB;
		break;

	case CASE(TINT8, TUINT8):
	case CASE(TUINT8, TUINT8):
	case CASE(TINT16, TUINT8):
	case CASE(TUINT16, TUINT8):
	case CASE(TINT32, TUINT8):
	case CASE(TUINT32, TUINT8):
		a = AMOVBU;
		break;

	case CASE(TINT64, TINT8):	// truncate low word
	case CASE(TUINT64, TINT8):
		a = AMOVB;
		goto trunc64;

	case CASE(TINT64, TUINT8):
	case CASE(TUINT64, TUINT8):
		a = AMOVBU;
		goto trunc64;

	case CASE(TINT16, TINT16):	// same size
	case CASE(TUINT16, TINT16):
	case CASE(TINT32, TINT16):	// truncate
	case CASE(TUINT32, TINT16):
		a = AMOVH;
		break;

	case CASE(TINT16, TUINT16):
	case CASE(TUINT16, TUINT16):
	case CASE(TINT32, TUINT16):
	case CASE(TUINT32, TUINT16):
		a = AMOVHU;
		break;

	case CASE(TINT64, TINT16):	// truncate low word
	case CASE(TUINT64, TINT16):
		a = AMOVH;
		goto trunc64;

	case CASE(TINT64, TUINT16):
	case CASE(TUINT64, TUINT16):
		a = AMOVHU;
		goto trunc64;

	case CASE(TINT32, TINT32):	// same size
	case CASE(TINT32, TUINT32):
	case CASE(TUINT32, TINT32):
	case CASE(TUINT32, TUINT32):
		a = AMOVW;
		break;

	case CASE(TINT64, TINT32):	// truncate
	case CASE(TUINT64, TINT32):
	case CASE(TINT64, TUINT32):
	case CASE(TUINT64, TUINT32):
		split64(f, &amp;flo, &amp;fhi);
		regalloc(&amp;r1, t-&gt;type, N);
		gins(AMOVW, &amp;flo, &amp;r1);
		gins(AMOVW, &amp;r1, t);
		regfree(&amp;r1);
		splitclean();
		return;

	case CASE(TINT64, TINT64):	// same size
	case CASE(TINT64, TUINT64):
	case CASE(TUINT64, TINT64):
	case CASE(TUINT64, TUINT64):
		split64(f, &amp;flo, &amp;fhi);
		split64(t, &amp;tlo, &amp;thi);
		regalloc(&amp;r1, flo.type, N);
		regalloc(&amp;r2, fhi.type, N);
		gins(AMOVW, &amp;flo, &amp;r1);
		gins(AMOVW, &amp;fhi, &amp;r2);
		gins(AMOVW, &amp;r1, &amp;tlo);
		gins(AMOVW, &amp;r2, &amp;thi);
		regfree(&amp;r1);
		regfree(&amp;r2);
		splitclean();
		splitclean();
		return;

	/*
	 * integer up-conversions
	 */
	case CASE(TINT8, TINT16):	// sign extend int8
	case CASE(TINT8, TUINT16):
	case CASE(TINT8, TINT32):
	case CASE(TINT8, TUINT32):
		a = AMOVB;
		goto rdst;
	case CASE(TINT8, TINT64):	// convert via int32
	case CASE(TINT8, TUINT64):
		cvt = types[TINT32];
		goto hard;

	case CASE(TUINT8, TINT16):	// zero extend uint8
	case CASE(TUINT8, TUINT16):
	case CASE(TUINT8, TINT32):
	case CASE(TUINT8, TUINT32):
		a = AMOVBU;
		goto rdst;
	case CASE(TUINT8, TINT64):	// convert via uint32
	case CASE(TUINT8, TUINT64):
		cvt = types[TUINT32];
		goto hard;

	case CASE(TINT16, TINT32):	// sign extend int16
	case CASE(TINT16, TUINT32):
		a = AMOVH;
		goto rdst;
	case CASE(TINT16, TINT64):	// convert via int32
	case CASE(TINT16, TUINT64):
		cvt = types[TINT32];
		goto hard;

	case CASE(TUINT16, TINT32):	// zero extend uint16
	case CASE(TUINT16, TUINT32):
		a = AMOVHU;
		goto rdst;
	case CASE(TUINT16, TINT64):	// convert via uint32
	case CASE(TUINT16, TUINT64):
		cvt = types[TUINT32];
		goto hard;

	case CASE(TINT32, TINT64):	// sign extend int32
	case CASE(TINT32, TUINT64):
		split64(t, &amp;tlo, &amp;thi);
		regalloc(&amp;r1, tlo.type, N);
		regalloc(&amp;r2, thi.type, N);
		gmove(f, &amp;r1);
		p1 = gins(AMOVW, &amp;r1, &amp;r2);
		p1-&gt;from.type = D_SHIFT;
		p1-&gt;from.offset = 2 &lt;&lt; 5 | 31 &lt;&lt; 7 | r1.val.u.reg; // r1-&gt;31
		p1-&gt;from.reg = NREG;
//print(&#34;gmove: %P\n&#34;, p1);
		gins(AMOVW, &amp;r1, &amp;tlo);
		gins(AMOVW, &amp;r2, &amp;thi);
		regfree(&amp;r1);
		regfree(&amp;r2);
		splitclean();
		return;

	case CASE(TUINT32, TINT64):	// zero extend uint32
	case CASE(TUINT32, TUINT64):
		split64(t, &amp;tlo, &amp;thi);
		gmove(f, &amp;tlo);
		regalloc(&amp;r1, thi.type, N);
		gins(AMOVW, ncon(0), &amp;r1);
		gins(AMOVW, &amp;r1, &amp;thi);
		regfree(&amp;r1);
		splitclean();
		return;

	/*
	* float to integer
	*/
	case CASE(TFLOAT32, TINT8):
	case CASE(TFLOAT32, TINT16):
	case CASE(TFLOAT32, TINT32):
	case CASE(TFLOAT32, TUINT8):
	case CASE(TFLOAT32, TUINT16):
	case CASE(TFLOAT32, TUINT32):
		fa = AMOVF;
		a = AMOVFW;
		ta = AMOVW;
		goto fltconv;

	case CASE(TFLOAT64, TINT8):
	case CASE(TFLOAT64, TINT16):
	case CASE(TFLOAT64, TINT32):
	case CASE(TFLOAT64, TUINT8):
	case CASE(TFLOAT64, TUINT16):
	case CASE(TFLOAT64, TUINT32):
		fa = AMOVD;
		a = AMOVDW;
		ta = AMOVW;
		goto fltconv;

	case CASE(TFLOAT32, TUINT64):
	case CASE(TFLOAT64, TUINT64):
		fatal(&#34;gmove TFLOAT, UINT64 not implemented&#34;);
		return;

	/*
	 * integer to float
	 */
	case CASE(TINT8, TFLOAT32):
	case CASE(TINT16, TFLOAT32):
	case CASE(TINT32, TFLOAT32):
	case CASE(TUINT8, TFLOAT32):
	case CASE(TUINT16, TFLOAT32):
	case CASE(TUINT32, TFLOAT32):
		fa = AMOVW;
		a = AMOVWF;
		ta = AMOVF;
		goto fltconv;

	case CASE(TINT8, TFLOAT64):
	case CASE(TINT16, TFLOAT64):
	case CASE(TINT32, TFLOAT64):
	case CASE(TUINT8, TFLOAT64):
	case CASE(TUINT16, TFLOAT64):
	case CASE(TUINT32, TFLOAT64):
		fa = AMOVW;
		a = AMOVWD;
		ta = AMOVD;
		goto fltconv;

	case CASE(TUINT64, TFLOAT32):
	case CASE(TUINT64, TFLOAT64):
		fatal(&#34;gmove UINT64, TFLOAT not implemented&#34;);
		return;


	/*
	 * float to float
	 */
	case CASE(TFLOAT32, TFLOAT32):
		a = AMOVF;
		break;

	case CASE(TFLOAT64, TFLOAT64):
		a = AMOVD;
		break;

	case CASE(TFLOAT32, TFLOAT64):
		regalloc(&amp;r1, types[TFLOAT64], t);
		gins(AMOVF, f, &amp;r1);
		gins(AMOVFD, &amp;r1, &amp;r1);
		gins(AMOVD, &amp;r1, t);
		regfree(&amp;r1);
		return;

	case CASE(TFLOAT64, TFLOAT32):
		regalloc(&amp;r1, types[TFLOAT64], t);
		gins(AMOVD, f, &amp;r1);
		gins(AMOVDF, &amp;r1, &amp;r1);
		gins(AMOVF, &amp;r1, t);
		regfree(&amp;r1);
		return;
	}

	gins(a, f, t);
	return;

rdst:
	// TODO(kaib): we almost always require a register dest anyway, this can probably be
	// removed.
	// requires register destination
	regalloc(&amp;r1, t-&gt;type, t);
	gins(a, f, &amp;r1);
	gmove(&amp;r1, t);
	regfree(&amp;r1);
	return;

hard:
	// requires register intermediate
	regalloc(&amp;r1, cvt, t);
	gmove(f, &amp;r1);
	gmove(&amp;r1, t);
	regfree(&amp;r1);
	return;

trunc64:
	// truncate 64 bit integer
	split64(f, &amp;flo, &amp;fhi);
	regalloc(&amp;r1, t-&gt;type, N);
	gins(a, &amp;flo, &amp;r1);
	gins(a, &amp;r1, t);
	regfree(&amp;r1);
	splitclean();
	return;

fltconv:
	regalloc(&amp;r1, types[ft], f);
	regalloc(&amp;r2, types[tt], t);
	gins(fa, f, &amp;r1);
	gins(a, &amp;r1, &amp;r2);
	gins(ta, &amp;r2, t);
	regfree(&amp;r1);
	regfree(&amp;r2);
	return;

fatal:
	// should not happen
	fatal(&#34;gmove %N -&gt; %N&#34;, f, t);
}

int
samaddr(Node *f, Node *t)
{

	if(f-&gt;op != t-&gt;op)
		return 0;

	switch(f-&gt;op) {
	case OREGISTER:
		if(f-&gt;val.u.reg != t-&gt;val.u.reg)
			break;
		return 1;
	}
	return 0;
}

/*
 * generate one instruction:
 *	as f, t
 */
Prog*
gins(int as, Node *f, Node *t)
{
//	Node nod;
//	int32 v;
	Prog *p;
	Addr af, at;

	if(f != N &amp;&amp; f-&gt;op == OINDEX) {
		fatal(&#34;gins OINDEX not implemented&#34;);
//		regalloc(&amp;nod, &amp;regnode, Z);
//		v = constnode.vconst;
//		cgen(f-&gt;right, &amp;nod);
//		constnode.vconst = v;
//		idx.reg = nod.reg;
//		regfree(&amp;nod);
	}
	if(t != N &amp;&amp; t-&gt;op == OINDEX) {
		fatal(&#34;gins OINDEX not implemented&#34;);
//		regalloc(&amp;nod, &amp;regnode, Z);
//		v = constnode.vconst;
//		cgen(t-&gt;right, &amp;nod);
//		constnode.vconst = v;
//		idx.reg = nod.reg;
//		regfree(&amp;nod);
	}

	memset(&amp;af, 0, sizeof af);
	memset(&amp;at, 0, sizeof at);
	if(f != N)
		naddr(f, &amp;af, 1);
	if(t != N)
		naddr(t, &amp;at, 1);	p = prog(as);
	if(f != N)
		p-&gt;from = af;
	if(t != N)
		p-&gt;to = at;
	if(debug[&#39;g&#39;])
		print(&#34;%P\n&#34;, p);
	return p;
}

/*
 * insert n into reg slot of p
 */
void
raddr(Node *n, Prog *p)
{
	Addr a;

	naddr(n, &amp;a, 1);
	if(a.type != D_REG &amp;&amp; a.type != D_FREG) {
		if(n)
			fatal(&#34;bad in raddr: %O&#34;, n-&gt;op);
		else
			fatal(&#34;bad in raddr: &lt;null&gt;&#34;);
		p-&gt;reg = NREG;
	} else
		p-&gt;reg = a.reg;
}

/* generate a comparison
TODO(kaib): one of the args can actually be a small constant. relax the constraint and fix call sites.
 */
Prog*
gcmp(int as, Node *lhs, Node *rhs)
{
	Prog *p;

	if(lhs-&gt;op != OREGISTER || rhs-&gt;op != OREGISTER)
		fatal(&#34;bad operands to gcmp: %O %O&#34;, lhs-&gt;op, rhs-&gt;op);

	p = gins(as, rhs, N);
	raddr(lhs, p);
	return p;
}

/* generate a constant shift
 * arm encodes a shift by 32 as 0, thus asking for 0 shift is illegal.
*/
Prog*
gshift(int as, Node *lhs, int32 stype, int32 sval, Node *rhs)
{
	Prog *p;

	if (sval &lt;= 0 || sval &gt; 32)
		fatal(&#34;bad shift value: %d&#34;, sval);

	sval = sval&amp;0x1f;

	p = gins(as, N, rhs);
	p-&gt;from.type = D_SHIFT;
	p-&gt;from.offset = stype | sval&lt;&lt;7 | lhs-&gt;val.u.reg;
	return p;
}

/* generate a register shift
*/
Prog *
gregshift(int as, Node *lhs, int32 stype, Node *reg, Node *rhs)
{
	Prog *p;
	p = gins(as, N, rhs);
	p-&gt;from.type = D_SHIFT;
	p-&gt;from.offset = stype | reg-&gt;val.u.reg &lt;&lt; 8 | 1&lt;&lt;4 | lhs-&gt;val.u.reg;
	return p;
}

static void
checkoffset(Addr *a, int canemitcode)
{
	Prog *p;
	Node n1;

	if(a-&gt;offset &lt; unmappedzero)
		return;
	if(!canemitcode)
		fatal(&#34;checkoffset %#llx, cannot emit code&#34;, a-&gt;offset);

	// cannot rely on unmapped nil page at 0 to catch
	// reference with large offset.  instead, emit explicit
	// test of 0(reg).
	regalloc(&amp;n1, types[TUINTPTR], N);
	p = gins(AMOVW, N, &amp;n1);
	p-&gt;from = *a;
	p-&gt;from.offset = 0;
	regfree(&amp;n1);
}

/*
 * generate code to compute n;
 * make a refer to result.
 */
void
naddr(Node *n, Addr *a, int canemitcode)
{
	a-&gt;type = D_NONE;
	a-&gt;name = D_NONE;
	a-&gt;reg = NREG;
	if(n == N)
		return;

	switch(n-&gt;op) {
	default:
		fatal(&#34;naddr: bad %O %D&#34;, n-&gt;op, a);
		break;

	case OREGISTER:
		if (n-&gt;val.u.reg &lt;= REGALLOC_RMAX) {
			a-&gt;type = D_REG;
			a-&gt;reg = n-&gt;val.u.reg;
		} else {
			a-&gt;type = D_FREG;
			a-&gt;reg = n-&gt;val.u.reg - REGALLOC_F0;
		}
		a-&gt;sym = S;
		break;

	case OINDEX:
	case OIND:
		fatal(&#34;naddr: OINDEX&#34;);
//		naddr(n-&gt;left, a);
//		if(a-&gt;type &gt;= D_AX &amp;&amp; a-&gt;type &lt;= D_DI)
//			a-&gt;type += D_INDIR;
//		else
//		if(a-&gt;type == D_CONST)
//			a-&gt;type = D_NONE+D_INDIR;
//		else
//		if(a-&gt;type == D_ADDR) {
//			a-&gt;type = a-&gt;index;
//			a-&gt;index = D_NONE;
//		} else
//			goto bad;
//		if(n-&gt;op == OINDEX) {
//			a-&gt;index = idx.reg;
//			a-&gt;scale = n-&gt;scale;
//		}
//		break;

	case OINDREG:
		a-&gt;type = D_OREG;
		a-&gt;reg = n-&gt;val.u.reg;
		a-&gt;sym = n-&gt;sym;
		a-&gt;offset = n-&gt;xoffset;
		checkoffset(a, canemitcode);
		break;

	case OPARAM:
		// n-&gt;left is PHEAP ONAME for stack parameter.
		// compute address of actual parameter on stack.
		a-&gt;etype = simtype[n-&gt;left-&gt;type-&gt;etype];
		a-&gt;width = n-&gt;left-&gt;type-&gt;width;
		a-&gt;offset = n-&gt;xoffset;
		a-&gt;sym = n-&gt;left-&gt;sym;
		a-&gt;type = D_OREG;
		a-&gt;name = D_PARAM;
		break;

	case ONAME:
		a-&gt;etype = 0;
		a-&gt;width = 0;
		a-&gt;reg = NREG;
		if(n-&gt;type != T) {
			a-&gt;etype = simtype[n-&gt;type-&gt;etype];
			a-&gt;width = n-&gt;type-&gt;width;
		}
		a-&gt;offset = n-&gt;xoffset;
		a-&gt;sym = n-&gt;sym;
		if(a-&gt;sym == S)
			a-&gt;sym = lookup(&#34;.noname&#34;);
		if(n-&gt;method) {
			if(n-&gt;type != T)
			if(n-&gt;type-&gt;sym != S)
			if(n-&gt;type-&gt;sym-&gt;package != nil)
				a-&gt;sym = pkglookup(a-&gt;sym-&gt;name, n-&gt;type-&gt;sym-&gt;package);
		}

		a-&gt;type = D_OREG;
		switch(n-&gt;class) {
		default:
			fatal(&#34;naddr: ONAME class %S %d\n&#34;, n-&gt;sym, n-&gt;class);
		case PEXTERN:
			a-&gt;name = D_EXTERN;
			break;
		case PAUTO:
			a-&gt;name = D_AUTO;
			break;
		case PPARAM:
		case PPARAMOUT:
			a-&gt;name = D_PARAM;
			break;
		case PFUNC:
			a-&gt;name = D_EXTERN;
			a-&gt;type = D_CONST;
			break;
		}
		break;

	case OLITERAL:
		switch(n-&gt;val.ctype) {
		default:
			fatal(&#34;naddr: const %lT&#34;, n-&gt;type);
			break;
		case CTFLT:
			a-&gt;type = D_FCONST;
			a-&gt;dval = mpgetflt(n-&gt;val.u.fval);
			break;
		case CTINT:
			a-&gt;sym = S;
			a-&gt;type = D_CONST;
			a-&gt;offset = mpgetfix(n-&gt;val.u.xval);
			break;
		case CTSTR:
			datagostring(n-&gt;val.u.sval, a);
			break;
		case CTBOOL:
			a-&gt;sym = S;
			a-&gt;type = D_CONST;
			a-&gt;offset = n-&gt;val.u.bval;
			break;
		case CTNIL:
			a-&gt;sym = S;
			a-&gt;type = D_CONST;
			a-&gt;offset = 0;
			break;
		}
		break;

	case OLEN:
		// len of string or slice
		naddr(n-&gt;left, a, canemitcode);
		a-&gt;offset += Array_nel;
		if(a-&gt;offset &gt;= unmappedzero &amp;&amp; a-&gt;offset-Array_nel &lt; unmappedzero)
			checkoffset(a, canemitcode);
		break;

	case OCAP:
		// cap of string or slice
		naddr(n-&gt;left, a, canemitcode);
		a-&gt;offset += Array_cap;
		if(a-&gt;offset &gt;= unmappedzero &amp;&amp; a-&gt;offset-Array_cap &lt; unmappedzero)
			checkoffset(a, canemitcode);
		break;

	case OADDR:
		naddr(n-&gt;left, a, canemitcode);
		switch(a-&gt;type) {
		case D_OREG:
			a-&gt;type = D_CONST;
			break;

		case D_REG:
		case D_CONST:
			break;
		
		default:
			fatal(&#34;naddr: OADDR %d\n&#34;, a-&gt;type);
		}
	}
}

/*
 * return Axxx for Oxxx on type t.
 */
int
optoas(int op, Type *t)
{
	int a;

	if(t == T)
		fatal(&#34;optoas: t is nil&#34;);

	a = AGOK;
	switch(CASE(op, simtype[t-&gt;etype])) {
	default:
		fatal(&#34;optoas: no entry %O-%T etype %T simtype %T&#34;, op, t, types[t-&gt;etype], types[simtype[t-&gt;etype]]);
		break;

/*	case CASE(OADDR, TPTR32):
		a = ALEAL;
		break;

	case CASE(OADDR, TPTR64):
		a = ALEAQ;
		break;
*/
	// TODO(kaib): make sure the conditional branches work on all edge cases
	case CASE(OEQ, TBOOL):
	case CASE(OEQ, TINT8):
	case CASE(OEQ, TUINT8):
	case CASE(OEQ, TINT16):
	case CASE(OEQ, TUINT16):
	case CASE(OEQ, TINT32):
	case CASE(OEQ, TUINT32):
	case CASE(OEQ, TINT64):
	case CASE(OEQ, TUINT64):
	case CASE(OEQ, TPTR32):
	case CASE(OEQ, TPTR64):
	case CASE(OEQ, TFLOAT32):
	case CASE(OEQ, TFLOAT64):
		a = ABEQ;
		break;

	case CASE(ONE, TBOOL):
	case CASE(ONE, TINT8):
	case CASE(ONE, TUINT8):
	case CASE(ONE, TINT16):
	case CASE(ONE, TUINT16):
	case CASE(ONE, TINT32):
	case CASE(ONE, TUINT32):
	case CASE(ONE, TINT64):
	case CASE(ONE, TUINT64):
	case CASE(ONE, TPTR32):
	case CASE(ONE, TPTR64):
	case CASE(ONE, TFLOAT32):
	case CASE(ONE, TFLOAT64):
		a = ABNE;
		break;

	case CASE(OLT, TINT8):
	case CASE(OLT, TINT16):
	case CASE(OLT, TINT32):
	case CASE(OLT, TINT64):
	case CASE(OLT, TFLOAT32):
	case CASE(OLT, TFLOAT64):
		a = ABLT;
		break;

	case CASE(OLT, TUINT8):
	case CASE(OLT, TUINT16):
	case CASE(OLT, TUINT32):
	case CASE(OLT, TUINT64):
		a = ABLO;
		break;

	case CASE(OLE, TINT8):
	case CASE(OLE, TINT16):
	case CASE(OLE, TINT32):
	case CASE(OLE, TINT64):
	case CASE(OLE, TFLOAT32):
	case CASE(OLE, TFLOAT64):
		a = ABLE;
		break;

	case CASE(OLE, TUINT8):
	case CASE(OLE, TUINT16):
	case CASE(OLE, TUINT32):
	case CASE(OLE, TUINT64):
		a = ABLS;
		break;

	case CASE(OGT, TINT8):
	case CASE(OGT, TINT16):
	case CASE(OGT, TINT32):
	case CASE(OGT, TINT64):
	case CASE(OGT, TFLOAT32):
	case CASE(OGT, TFLOAT64):
		a = ABGT;
		break;

	case CASE(OGT, TUINT8):
	case CASE(OGT, TUINT16):
	case CASE(OGT, TUINT32):
	case CASE(OGT, TUINT64):
		a = ABHI;
		break;

	case CASE(OGE, TINT8):
	case CASE(OGE, TINT16):
	case CASE(OGE, TINT32):
	case CASE(OGE, TINT64):
	case CASE(OGE, TFLOAT32):
	case CASE(OGE, TFLOAT64):
		a = ABGE;
		break;

	case CASE(OGE, TUINT8):
	case CASE(OGE, TUINT16):
	case CASE(OGE, TUINT32):
	case CASE(OGE, TUINT64):
		a = ABHS;
		break;

	case CASE(OCMP, TBOOL):
	case CASE(OCMP, TINT8):
	case CASE(OCMP, TUINT8):
	case CASE(OCMP, TINT16):
	case CASE(OCMP, TUINT16):
	case CASE(OCMP, TINT32):
	case CASE(OCMP, TUINT32):
	case CASE(OCMP, TPTR32):
		a = ACMP;
		break;

	case CASE(OCMP, TFLOAT32):
		a = ACMPF;
		break;

	case CASE(OCMP, TFLOAT64):
		a = ACMPD;
		break;

	case CASE(OAS, TBOOL):
	case CASE(OAS, TINT8):
	case CASE(OAS, TUINT8):
		a = AMOVB;
		break;

	case CASE(OAS, TINT16):
	case CASE(OAS, TUINT16):
		a = AMOVH;
		break;

	case CASE(OAS, TINT32):
	case CASE(OAS, TUINT32):
	case CASE(OAS, TPTR32):
		a = AMOVW;
		break;

	case CASE(OAS, TFLOAT32):
		a = AMOVF;
		break;

	case CASE(OAS, TFLOAT64):
		a = AMOVD;
		break;

	case CASE(OADD, TINT8):
	case CASE(OADD, TUINT8):
	case CASE(OADD, TINT16):
	case CASE(OADD, TUINT16):
	case CASE(OADD, TINT32):
	case CASE(OADD, TUINT32):
	case CASE(OADD, TPTR32):
		a = AADD;
		break;

	case CASE(OADD, TFLOAT32):
		a = AADDF;
		break;

	case CASE(OADD, TFLOAT64):
		a = AADDD;
		break;

	case CASE(OSUB, TINT8):
	case CASE(OSUB, TUINT8):
	case CASE(OSUB, TINT16):
	case CASE(OSUB, TUINT16):
	case CASE(OSUB, TINT32):
	case CASE(OSUB, TUINT32):
	case CASE(OSUB, TPTR32):
		a = ASUB;
		break;

	case CASE(OSUB, TFLOAT32):
		a = ASUBF;
		break;

	case CASE(OSUB, TFLOAT64):
		a = ASUBD;
		break;

	case CASE(OAND, TINT8):
	case CASE(OAND, TUINT8):
	case CASE(OAND, TINT16):
	case CASE(OAND, TUINT16):
	case CASE(OAND, TINT32):
	case CASE(OAND, TUINT32):
	case CASE(OAND, TPTR32):
		a = AAND;
		break;

	case CASE(OOR, TINT8):
	case CASE(OOR, TUINT8):
	case CASE(OOR, TINT16):
	case CASE(OOR, TUINT16):
	case CASE(OOR, TINT32):
	case CASE(OOR, TUINT32):
	case CASE(OOR, TPTR32):
		a = AORR;
		break;

	case CASE(OXOR, TINT8):
	case CASE(OXOR, TUINT8):
	case CASE(OXOR, TINT16):
	case CASE(OXOR, TUINT16):
	case CASE(OXOR, TINT32):
	case CASE(OXOR, TUINT32):
	case CASE(OXOR, TPTR32):
		a = AEOR;
		break;

	case CASE(OLSH, TINT8):
	case CASE(OLSH, TUINT8):
	case CASE(OLSH, TINT16):
	case CASE(OLSH, TUINT16):
	case CASE(OLSH, TINT32):
	case CASE(OLSH, TUINT32):
	case CASE(OLSH, TPTR32):
		a = ASLL;
		break;

	case CASE(ORSH, TUINT8):
	case CASE(ORSH, TUINT16):
	case CASE(ORSH, TUINT32):
	case CASE(ORSH, TPTR32):
		a = ASRL;
		break;

	case CASE(ORSH, TINT8):
	case CASE(ORSH, TINT16):
	case CASE(ORSH, TINT32):
		a = ASRA;
		break;

	case CASE(OMUL, TUINT8):
	case CASE(OMUL, TUINT16):
	case CASE(OMUL, TUINT32):
	case CASE(OMUL, TPTR32):
		a = AMULU;
		break;

	case CASE(OMUL, TINT8):
	case CASE(OMUL, TINT16):
	case CASE(OMUL, TINT32):
		a = AMUL;
		break;

	case CASE(OMUL, TFLOAT32):
		a = AMULF;
		break;

	case CASE(OMUL, TFLOAT64):
		a = AMULD;
		break;

	case CASE(ODIV, TUINT8):
	case CASE(ODIV, TUINT16):
	case CASE(ODIV, TUINT32):
	case CASE(ODIV, TPTR32):
		a = ADIVU;
		break;

	case CASE(ODIV, TINT8):
	case CASE(ODIV, TINT16):
	case CASE(ODIV, TINT32):
		a = ADIV;
		break;

	case CASE(OMOD, TUINT8):
	case CASE(OMOD, TUINT16):
	case CASE(OMOD, TUINT32):
	case CASE(OMOD, TPTR32):
		a = AMODU;
		break;

	case CASE(OMOD, TINT8):
	case CASE(OMOD, TINT16):
	case CASE(OMOD, TINT32):
		a = AMOD;
		break;

//	case CASE(OEXTEND, TINT16):
//		a = ACWD;
//		break;

//	case CASE(OEXTEND, TINT32):
//		a = ACDQ;
//		break;

//	case CASE(OEXTEND, TINT64):
//		a = ACQO;
//		break;

	case CASE(ODIV, TFLOAT32):
		a = ADIVF;
		break;

	case CASE(ODIV, TFLOAT64):
		a = ADIVD;
		break;

	}
	return a;
}

enum
{
	ODynam	= 1&lt;&lt;0,
	OPtrto	= 1&lt;&lt;1,
};

static	Node	clean[20];
static	int	cleani = 0;

void
sudoclean(void)
{
	if(clean[cleani-1].op != OEMPTY)
		regfree(&amp;clean[cleani-1]);
	if(clean[cleani-2].op != OEMPTY)
		regfree(&amp;clean[cleani-2]);
	cleani -= 2;
}

int
dotaddable(Node *n, Node *n1)
{
	int o, oary[10];
	Node *nn;

	if(n-&gt;op != ODOT)
		return 0;

	o = dotoffset(n, oary, &amp;nn);
	if(nn != N &amp;&amp; nn-&gt;addable &amp;&amp; o == 1 &amp;&amp; oary[0] &gt;= 0) {
		*n1 = *nn;
		n1-&gt;type = n-&gt;type;
		n1-&gt;xoffset += oary[0];
		return 1;
	}
	return 0;
}

/*
 * generate code to compute address of n,
 * a reference to a (perhaps nested) field inside
 * an array or struct.
 * return 0 on failure, 1 on success.
 * on success, leaves usable address in a.
 *
 * caller is responsible for calling sudoclean
 * after successful sudoaddable,
 * to release the register used for a.
 */
int
sudoaddable(int as, Node *n, Addr *a, int *w)
{
	int o, i;
	int oary[10];
	int64 v;
	Node n1, n2, n3, n4, *nn, *l, *r;
	Node *reg, *reg1;
	Prog *p1;
	Type *t;

	if(n-&gt;type == T)
		return 0;

	switch(n-&gt;op) {
	case OLITERAL:
		if(n-&gt;val.ctype != CTINT)
			break;
		v = mpgetfix(n-&gt;val.u.xval);
		if(v &gt;= 32000 || v &lt;= -32000)
			break;
		goto lit;

	case ODOT:
	case ODOTPTR:
		cleani += 2;
		reg = &amp;clean[cleani-1];
		reg1 = &amp;clean[cleani-2];
		reg-&gt;op = OEMPTY;
		reg1-&gt;op = OEMPTY;
		goto odot;

	case OINDEX:
		cleani += 2;
		reg = &amp;clean[cleani-1];
		reg1 = &amp;clean[cleani-2];
		reg-&gt;op = OEMPTY;
		reg1-&gt;op = OEMPTY;
		goto oindex;
	}
	return 0;

lit:
	switch(as) {
	default:
		return 0;
	case AADD: case ASUB: case AAND: case AORR: case AEOR:
	case AMOVB: case AMOVBU: case AMOVH: case AMOVHU:
	case AMOVW:
		break;
	}

	cleani += 2;
	reg = &amp;clean[cleani-1];
	reg1 = &amp;clean[cleani-2];
	reg-&gt;op = OEMPTY;
	reg1-&gt;op = OEMPTY;
	naddr(n, a, 1);
	goto yes;

odot:
	o = dotoffset(n, oary, &amp;nn);
	if(nn == N)
		goto no;

	if(nn-&gt;addable &amp;&amp; o == 1 &amp;&amp; oary[0] &gt;= 0) {
		// directly addressable set of DOTs
		n1 = *nn;
		n1.type = n-&gt;type;
		n1.xoffset += oary[0];
		naddr(&amp;n1, a, 1);
		goto yes;
	}

	regalloc(reg, types[tptr], N);
	n1 = *reg;
	n1.op = OINDREG;
	if(oary[0] &gt;= 0) {
		agen(nn, reg);
		n1.xoffset = oary[0];
	} else {
		cgen(nn, reg);
		n1.xoffset = -(oary[0]+1);
	}

	for(i=1; i&lt;o; i++) {
		if(oary[i] &gt;= 0)
			fatal(&#34;cant happen&#34;);
		gins(AMOVW, &amp;n1, reg);
		n1.xoffset = -(oary[i]+1);
	}

	a-&gt;type = D_NONE;
	a-&gt;name = D_NONE;
	naddr(&amp;n1, a, 1);
	goto yes;

oindex:
	l = n-&gt;left;
	r = n-&gt;right;
	if(l-&gt;ullman &gt;= UINF &amp;&amp; r-&gt;ullman &gt;= UINF)
		goto no;

	// set o to type of array
	o = 0;
	if(isptr[l-&gt;type-&gt;etype]) {
		o += OPtrto;
		if(l-&gt;type-&gt;type-&gt;etype != TARRAY)
			fatal(&#34;not ptr ary&#34;);
		if(l-&gt;type-&gt;type-&gt;bound &lt; 0)
			o += ODynam;
	} else {
		if(l-&gt;type-&gt;etype != TARRAY)
			fatal(&#34;not ary&#34;);
		if(l-&gt;type-&gt;bound &lt; 0)
			o += ODynam;
	}

	*w = n-&gt;type-&gt;width;
	if(isconst(r, CTINT))
		goto oindex_const;

	switch(*w) {
	default:
		goto no;
	case 1:
	case 2:
	case 4:
	case 8:
		break;
	}

	// load the array (reg)
	if(l-&gt;ullman &gt; r-&gt;ullman) {
		regalloc(reg, types[tptr], N);
		if(o &amp; OPtrto)
			cgen(l, reg);
		else
			agen(l, reg);
	}

	// load the index (reg1)
	t = types[TUINT32];
	if(issigned[r-&gt;type-&gt;etype])
		t = types[TINT32];
	regalloc(reg1, t, N);
	regalloc(&amp;n3, r-&gt;type, reg1);
	cgen(r, &amp;n3);
	gmove(&amp;n3, reg1);
	regfree(&amp;n3);

	// load the array (reg)
	if(l-&gt;ullman &lt;= r-&gt;ullman) {
		regalloc(reg, types[tptr], N);
		if(o &amp; OPtrto)
			cgen(l, reg);
		else
			agen(l, reg);
	}

	// check bounds
	if(!debug[&#39;B&#39;]) {
		if(o &amp; ODynam) {
			n2 = *reg;
			n2.op = OINDREG;
			n2.type = types[tptr];
			n2.xoffset = Array_nel;
		} else {
			if(l-&gt;type-&gt;width &gt;= unmappedzero &amp;&amp; l-&gt;op == OIND) {
				// cannot rely on page protections to
				// catch array ptr == 0, so dereference.
				n2 = *reg;
				n2.op = OINDREG;
				n2.type = types[TUINTPTR];
				n2.xoffset = 0;
				regalloc(&amp;n3, n2.type, N);
				gins(AMOVW, &amp;n2, &amp;n3);
				regfree(&amp;n3);
			}
			nodconst(&amp;n2, types[TUINT32], l-&gt;type-&gt;bound);
			if(o &amp; OPtrto)
				nodconst(&amp;n2, types[TUINT32], l-&gt;type-&gt;type-&gt;bound);
		}
		regalloc(&amp;n3, n2.type, N);
		cgen(&amp;n2, &amp;n3);
		gcmp(optoas(OCMP, types[TUINT32]), reg1, &amp;n3);
		regfree(&amp;n3);
		p1 = gbranch(optoas(OLT, types[TUINT32]), T);
		ginscall(throwindex, 0);
		patch(p1, pc);
	}

	if(o &amp; ODynam) {
		n2 = *reg;
		n2.op = OINDREG;
		n2.type = types[tptr];
		n2.xoffset = Array_array;
		gmove(&amp;n2, reg);
	}

	if (*w == 1)
		gins(AADD, reg1, reg);
	else if(*w == 2)
		gshift(AADD, reg1, SHIFT_LL, 1, reg);
	else if(*w == 4)
		gshift(AADD, reg1, SHIFT_LL, 2, reg);
	else if(*w == 8)
		gshift(AADD, reg1, SHIFT_LL, 3, reg);

	naddr(reg1, a, 1);
	a-&gt;type = D_OREG;
	a-&gt;reg = reg-&gt;val.u.reg;
	a-&gt;offset = 0;

	goto yes;

oindex_const:
	// index is constant
	// can check statically and
	// can multiply by width statically

	regalloc(reg, types[tptr], N);
	if(o &amp; OPtrto)
		cgen(l, reg);
	else
		agen(l, reg);

	v = mpgetfix(r-&gt;val.u.xval);
	if(o &amp; ODynam) {

		if(!debug[&#39;B&#39;]) {
			n1 = *reg;
			n1.op = OINDREG;
			n1.type = types[tptr];
			n1.xoffset = Array_nel;
			nodconst(&amp;n2, types[TUINT32], v);
			regalloc(&amp;n3, types[TUINT32], N);
			cgen(&amp;n2, &amp;n3);
			regalloc(&amp;n4, n1.type, N);
			cgen(&amp;n1, &amp;n4);
			gcmp(optoas(OCMP, types[TUINT32]), &amp;n4, &amp;n3);
			regfree(&amp;n4);
			regfree(&amp;n3);
			p1 = gbranch(optoas(OGT, types[TUINT32]), T);
			ginscall(throwindex, 0);
			patch(p1, pc);
		}

		n1 = *reg;
		n1.op = OINDREG;
		n1.type = types[tptr];
		n1.xoffset = Array_array;
		gmove(&amp;n1, reg);

	} else
	if(!debug[&#39;B&#39;]) {
		if(v &lt; 0) {
			yyerror(&#34;out of bounds on array&#34;);
		} else
		if(o &amp; OPtrto) {
			if(v &gt;= l-&gt;type-&gt;type-&gt;bound)
				yyerror(&#34;out of bounds on array&#34;);
		} else
		if(v &gt;= l-&gt;type-&gt;bound) {
			yyerror(&#34;out of bounds on array&#34;);
		}
	}

	n2 = *reg;
	n2.op = OINDREG;
	n2.xoffset = v * (*w);
	a-&gt;type = D_NONE;
	a-&gt;name = D_NONE;
	naddr(&amp;n2, a, 1);
	goto yes;

yes:
	return 1;

no:
	sudoclean();
	return 0;
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
