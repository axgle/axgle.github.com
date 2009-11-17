<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/align.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/align.c</h1>

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
 * machine size and rounding
 * alignment is dictated around
 * the size of a pointer, set in belexinit
 * (see ../6g/align.c).
 */

static int defercalc;

uint32
rnd(uint32 o, uint32 r)
{
	if(maxround == 0)
		fatal(&#34;rnd&#34;);

	if(r &gt; maxround)
		r = maxround;
	if(r != 0)
		while(o%r != 0)
			o++;
	return o;
}

static void
offmod(Type *t)
{
	Type *f;
	int32 o;

	o = 0;
	for(f=t-&gt;type; f!=T; f=f-&gt;down) {
		if(f-&gt;etype != TFIELD)
			fatal(&#34;widstruct: not TFIELD: %lT&#34;, f);
		if(f-&gt;type-&gt;etype != TFUNC)
			continue;
		f-&gt;width = o;
		o += widthptr;
	}
}

static uint32
arrayelemwidth(Type *t)
{

	while(t-&gt;etype == TARRAY &amp;&amp; t-&gt;bound &gt;= 0)
		t = t-&gt;type;
	return t-&gt;width;
}

static uint32
widstruct(Type *t, uint32 o, int flag)
{
	Type *f;
	int32 w, m;

	for(f=t-&gt;type; f!=T; f=f-&gt;down) {
		if(f-&gt;etype != TFIELD)
			fatal(&#34;widstruct: not TFIELD: %lT&#34;, f);
		dowidth(f-&gt;type);
		if(f-&gt;type-&gt;width &lt; 0)
			fatal(&#34;invalid width %lld&#34;, f-&gt;type-&gt;width);
		w = f-&gt;type-&gt;width;
		m = arrayelemwidth(f-&gt;type);
		o = rnd(o, m);
		f-&gt;width = o;	// really offset for TFIELD
		if(f-&gt;nname != N) {
			// this same stackparam logic is in addrescapes
			// in typecheck.c.  usually addrescapes runs after
			// widstruct, in which case we could drop this,
			// but function closure functions are the exception.
			if(f-&gt;nname-&gt;stackparam) {
				f-&gt;nname-&gt;stackparam-&gt;xoffset = o;
				f-&gt;nname-&gt;xoffset = 0;
			} else
				f-&gt;nname-&gt;xoffset = o;
		}
		o += w;
	}
	// final width is rounded
	if(flag)
		o = rnd(o, maxround);

	// type width only includes back to first field&#39;s offset
	if(t-&gt;type == T)
		t-&gt;width = 0;
	else
		t-&gt;width = o - t-&gt;type-&gt;width;
	return o;
}

void
dowidth(Type *t)
{
	int32 et;
	uint32 w;
	int lno;
	Type *t1;

	if(maxround == 0 || widthptr == 0)
		fatal(&#34;dowidth without betypeinit&#34;);

	if(t == T)
		return;

	if(t-&gt;width &gt; 0)
		return;

	if(t-&gt;width == -2) {
		lno = lineno;
		lineno = t-&gt;lineno;
		yyerror(&#34;invalid recursive type %T&#34;, t);
		t-&gt;width = 0;
		lineno = lno;
		return;
	}

	// defer checkwidth calls until after we&#39;re done
	defercalc++;

	lno = lineno;
	lineno = t-&gt;lineno;
	t-&gt;width = -2;

	et = t-&gt;etype;
	switch(et) {
	case TFUNC:
	case TCHAN:
	case TMAP:
	case TSTRING:
		break;

	default:
		/* simtype == 0 during bootstrap */
		if(simtype[t-&gt;etype] != 0)
			et = simtype[t-&gt;etype];
		break;
	}

	w = 0;
	switch(et) {
	default:
		fatal(&#34;dowidth: unknown type: %T&#34;, t);
		break;

	/* compiler-specific stuff */
	case TINT8:
	case TUINT8:
	case TBOOL:		// bool is int8
		w = 1;
		break;
	case TINT16:
	case TUINT16:
		w = 2;
		break;
	case TINT32:
	case TUINT32:
	case TFLOAT32:
		w = 4;
		break;
	case TINT64:
	case TUINT64:
	case TFLOAT64:
		w = 8;
		break;
	case TPTR32:
		w = 4;
		checkwidth(t-&gt;type);
		break;
	case TPTR64:
		w = 8;
		checkwidth(t-&gt;type);
		break;
	case TDDD:
		w = 2*widthptr;
		break;
	case TINTER:		// implemented as 2 pointers
		w = 2*widthptr;
		offmod(t);
		break;
	case TCHAN:		// implemented as pointer
		w = widthptr;
		checkwidth(t-&gt;type);
		break;
	case TMAP:		// implemented as pointer
		w = widthptr;
		checkwidth(t-&gt;type);
		checkwidth(t-&gt;down);
		break;
	case TFORW:		// should have been filled in
	case TANY:
		// dummy type; should be replaced before use.
		if(!debug[&#39;A&#39;])
			fatal(&#34;dowidth any&#34;);
		w = 1;	// anything will do
		break;
	case TSTRING:
		if(sizeof_String == 0)
			fatal(&#34;early dowidth string&#34;);
		w = sizeof_String;
		break;
	case TARRAY:
		if(t-&gt;type == T)
			break;
		if(t-&gt;bound &gt;= 0) {
			dowidth(t-&gt;type);
			w = t-&gt;bound * t-&gt;type-&gt;width;
			if(w == 0)
				w = maxround;
		}
		else if(t-&gt;bound == -1) {
			w = sizeof_Array;
			checkwidth(t-&gt;type);
		}
		else
			fatal(&#34;dowidth %T&#34;, t);	// probably [...]T
		break;

	case TSTRUCT:
		if(t-&gt;funarg)
			fatal(&#34;dowidth fn struct %T&#34;, t);
		w = widstruct(t, 0, 1);
		if(w == 0)
			w = maxround;
		break;

	case TFUNC:
		// make fake type to check later to
		// trigger function argument computation.
		t1 = typ(TFUNCARGS);
		t1-&gt;type = t;
		checkwidth(t1);

		// width of func type is pointer
		w = widthptr;
		break;

	case TFUNCARGS:
		// function is 3 cated structures;
		// compute their widths as side-effect.
		t1 = t-&gt;type;
		w = widstruct(*getthis(t1), 0, 0);
		w = widstruct(*getinarg(t1), w, 1);
		w = widstruct(*getoutarg(t1), w, 1);
		t1-&gt;argwid = w;
		break;
	}

	t-&gt;width = w;
	lineno = lno;

	if(defercalc == 1)
		resumecheckwidth();
	else
		--defercalc;
}

/*
 * when a type&#39;s width should be known, we call checkwidth
 * to compute it.  during a declaration like
 *
 *	type T *struct { next T }
 *
 * it is necessary to defer the calculation of the struct width
 * until after T has been initialized to be a pointer to that struct.
 * similarly, during import processing structs may be used
 * before their definition.  in those situations, calling
 * defercheckwidth() stops width calculations until
 * resumecheckwidth() is called, at which point all the
 * checkwidths that were deferred are executed.
 * dowidth should only be called when the type&#39;s size
 * is needed immediately.  checkwidth makes sure the
 * size is evaluated eventually.
 */
typedef struct TypeList TypeList;
struct TypeList {
	Type *t;
	TypeList *next;
};

static TypeList *tlfree;
static TypeList *tlq;

void
checkwidth(Type *t)
{
	TypeList *l;

	if(t == T)
		return;

	// function arg structs should not be checked
	// outside of the enclosing function.
	if(t-&gt;funarg)
		fatal(&#34;checkwidth %T&#34;, t);

	if(!defercalc) {
		dowidth(t);
		return;
	}
	if(t-&gt;deferwidth)
		return;
	t-&gt;deferwidth = 1;

	l = tlfree;
	if(l != nil)
		tlfree = l-&gt;next;
	else
		l = mal(sizeof *l);

	l-&gt;t = t;
	l-&gt;next = tlq;
	tlq = l;
}

void
defercheckwidth(void)
{
	// we get out of sync on syntax errors, so don&#39;t be pedantic.
	// if(defercalc)
	//	fatal(&#34;defercheckwidth&#34;);
	defercalc = 1;
}

void
resumecheckwidth(void)
{
	TypeList *l;

	if(!defercalc)
		fatal(&#34;resumecheckwidth&#34;);
	for(l = tlq; l != nil; l = tlq) {
		l-&gt;t-&gt;deferwidth = 0;
		tlq = l-&gt;next;
		dowidth(l-&gt;t);
		l-&gt;next = tlfree;
		tlfree = l;
	}
	defercalc = 0;
}

void
typeinit(void)
{
	int i, etype, sameas;
	Type *t;
	Sym *s, *s1;

	if(widthptr == 0)
		fatal(&#34;typeinit before betypeinit&#34;);

	for(i=0; i&lt;NTYPE; i++)
		simtype[i] = i;

	types[TPTR32] = typ(TPTR32);
	dowidth(types[TPTR32]);

	types[TPTR64] = typ(TPTR64);
	dowidth(types[TPTR64]);

	tptr = TPTR32;
	if(widthptr == 8)
		tptr = TPTR64;

	for(i=TINT8; i&lt;=TUINT64; i++)
		isint[i] = 1;
	isint[TINT] = 1;
	isint[TUINT] = 1;
	isint[TUINTPTR] = 1;

	for(i=TFLOAT32; i&lt;=TFLOAT64; i++)
		isfloat[i] = 1;
	isfloat[TFLOAT] = 1;

	isptr[TPTR32] = 1;
	isptr[TPTR64] = 1;

	isforw[TFORW] = 1;

	issigned[TINT] = 1;
	issigned[TINT8] = 1;
	issigned[TINT16] = 1;
	issigned[TINT32] = 1;
	issigned[TINT64] = 1;

	/*
	 * initialize okfor
	 */
	for(i=0; i&lt;NTYPE; i++) {
		if(isint[i] || i == TIDEAL) {
			okforeq[i] = 1;
			okforcmp[i] = 1;
			okforarith[i] = 1;
			okforadd[i] = 1;
			okforand[i] = 1;
			issimple[i] = 1;
			minintval[i] = mal(sizeof(*minintval[i]));
			maxintval[i] = mal(sizeof(*maxintval[i]));
		}
		if(isfloat[i]) {
			okforeq[i] = 1;
			okforcmp[i] = 1;
			okforadd[i] = 1;
			okforarith[i] = 1;
			issimple[i] = 1;
			minfltval[i] = mal(sizeof(*minfltval[i]));
			maxfltval[i] = mal(sizeof(*maxfltval[i]));
		}
	}

	issimple[TBOOL] = 1;

	okforadd[TSTRING] = 1;

	okforbool[TBOOL] = 1;

	okforcap[TARRAY] = 1;
	okforcap[TCHAN] = 1;

	okforlen[TARRAY] = 1;
	okforlen[TCHAN] = 1;
	okforlen[TMAP] = 1;
	okforlen[TSTRING] = 1;

	okforeq[TPTR32] = 1;
	okforeq[TPTR64] = 1;
	okforeq[TINTER] = 1;
	okforeq[TMAP] = 1;
	okforeq[TCHAN] = 1;
	okforeq[TFUNC] = 1;
	okforeq[TSTRING] = 1;
	okforeq[TBOOL] = 1;
	okforeq[TARRAY] = 1;	// refined in typecheck

	okforcmp[TSTRING] = 1;

	for(i=0; i&lt;nelem(okfor); i++)
		okfor[i] = okfornone;

	// binary
	okfor[OADD] = okforadd;
	okfor[OAND] = okforand;
	okfor[OANDAND] = okforbool;
	okfor[OANDNOT] = okforand;
	okfor[ODIV] = okforarith;
	okfor[OEQ] = okforeq;
	okfor[OGE] = okforcmp;
	okfor[OGT] = okforcmp;
	okfor[OLE] = okforcmp;
	okfor[OLT] = okforcmp;
	okfor[OMOD] = okforarith;
	okfor[OMUL] = okforarith;
	okfor[ONE] = okforeq;
	okfor[OOR] = okforand;
	okfor[OOROR] = okforbool;
	okfor[OSUB] = okforarith;
	okfor[OXOR] = okforand;
	okfor[OLSH] = okforand;
	okfor[ORSH] = okforand;

	// unary
	okfor[OCOM] = okforand;
	okfor[OMINUS] = okforarith;
	okfor[ONOT] = okforbool;
	okfor[OPLUS] = okforadd;

	// special
	okfor[OCAP] = okforcap;
	okfor[OLEN] = okforlen;

	// comparison
	iscmp[OLT] = 1;
	iscmp[OGT] = 1;
	iscmp[OGE] = 1;
	iscmp[OLE] = 1;
	iscmp[OEQ] = 1;
	iscmp[ONE] = 1;

	mpatofix(maxintval[TINT8], &#34;0x7f&#34;);
	mpatofix(minintval[TINT8], &#34;-0x80&#34;);
	mpatofix(maxintval[TINT16], &#34;0x7fff&#34;);
	mpatofix(minintval[TINT16], &#34;-0x8000&#34;);
	mpatofix(maxintval[TINT32], &#34;0x7fffffff&#34;);
	mpatofix(minintval[TINT32], &#34;-0x80000000&#34;);
	mpatofix(maxintval[TINT64], &#34;0x7fffffffffffffff&#34;);
	mpatofix(minintval[TINT64], &#34;-0x8000000000000000&#34;);

	mpatofix(maxintval[TUINT8], &#34;0xff&#34;);
	mpatofix(maxintval[TUINT16], &#34;0xffff&#34;);
	mpatofix(maxintval[TUINT32], &#34;0xffffffff&#34;);
	mpatofix(maxintval[TUINT64], &#34;0xffffffffffffffff&#34;);

	mpatoflt(maxfltval[TFLOAT32], &#34;3.40282347e+38&#34;);
	mpatoflt(minfltval[TFLOAT32], &#34;-3.40282347e+38&#34;);
	mpatoflt(maxfltval[TFLOAT64], &#34;1.7976931348623157e+308&#34;);
	mpatoflt(minfltval[TFLOAT64], &#34;-1.7976931348623157e+308&#34;);

	/* for walk to use in error messages */
	types[TFUNC] = functype(N, nil, nil);

	/* types used in front end */
	// types[TNIL] got set early in lexinit
	types[TIDEAL] = typ(TIDEAL);

	/* simple aliases */
	simtype[TMAP] = tptr;
	simtype[TCHAN] = tptr;
	simtype[TFUNC] = tptr;

	/* pick up the backend typedefs */
	for(i=0; typedefs[i].name; i++) {
		s = lookup(typedefs[i].name);
		s1 = pkglookup(typedefs[i].name, &#34;/builtin/&#34;);

		etype = typedefs[i].etype;
		if(etype &lt; 0 || etype &gt;= nelem(types))
			fatal(&#34;typeinit: %s bad etype&#34;, s-&gt;name);
		sameas = typedefs[i].sameas;
		if(sameas &lt; 0 || sameas &gt;= nelem(types))
			fatal(&#34;typeinit: %s bad sameas&#34;, s-&gt;name);
		simtype[etype] = sameas;
		minfltval[etype] = minfltval[sameas];
		maxfltval[etype] = maxfltval[sameas];
		minintval[etype] = minintval[sameas];
		maxintval[etype] = maxintval[sameas];

		t = types[etype];
		if(t != T)
			fatal(&#34;typeinit: %s already defined&#34;, s-&gt;name);

		t = typ(etype);
		t-&gt;sym = s;

		dowidth(t);
		types[etype] = t;
		s1-&gt;def = typenod(t);
	}

	Array_array = rnd(0, widthptr);
	Array_nel = rnd(Array_array+widthptr, types[TUINT32]-&gt;width);
	Array_cap = rnd(Array_nel+types[TUINT32]-&gt;width, types[TUINT32]-&gt;width);
	sizeof_Array = rnd(Array_cap+types[TUINT32]-&gt;width, maxround);

	// string is same as slice wo the cap
	sizeof_String = rnd(Array_nel+types[TUINT32]-&gt;width, maxround);

	dowidth(types[TSTRING]);
	dowidth(idealstring);
}

/*
 * compute total size of f&#39;s in/out arguments.
 */
int
argsize(Type *t)
{
	Iter save;
	Type *fp;
	int w, x;

	w = 0;

	fp = structfirst(&amp;save, getoutarg(t));
	while(fp != T) {
		x = fp-&gt;width + fp-&gt;type-&gt;width;
		if(x &gt; w)
			w = x;
		fp = structnext(&amp;save);
	}

	fp = funcfirst(&amp;save, t);
	while(fp != T) {
		x = fp-&gt;width + fp-&gt;type-&gt;width;
		if(x &gt; w)
			w = x;
		fp = funcnext(&amp;save);
	}

	w = (w+7) &amp; ~7;
	return w;
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
