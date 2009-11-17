<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/lib9/fmt/dofmt.c</title>

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
  <h1 id="generatedHeader">Text file src/lib9/fmt/dofmt.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
/*
 * The authors of this software are Rob Pike and Ken Thompson,
 * with contributions from Mike Burrows and Sean Dorward.
 *
 *     Copyright (c) 2002-2006 by Lucent Technologies.
 *     Portions Copyright (c) 2004 Google Inc.
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose without fee is hereby granted, provided that this entire notice
 * is included in all copies of any software which is or includes a copy
 * or modification of this software and in all copies of the supporting
 * documentation for such software.
 * THIS SOFTWARE IS BEING PROVIDED &#34;AS IS&#34;, WITHOUT ANY EXPRESS OR IMPLIED
 * WARRANTY.  IN PARTICULAR, NEITHER THE AUTHORS NOR LUCENT TECHNOLOGIES
 * NOR GOOGLE INC MAKE ANY REPRESENTATION OR WARRANTY OF ANY KIND CONCERNING
 * THE MERCHANTABILITY OF THIS SOFTWARE OR ITS FITNESS FOR ANY PARTICULAR PURPOSE.
 */

#include &lt;u.h&gt;
#include &lt;libc.h&gt;
#include &#34;fmtdef.h&#34;

/* format the output into f-&gt;to and return the number of characters fmted  */
int
dofmt(Fmt *f, char *fmt)
{
	Rune rune, *rt, *rs;
	int r;
	char *t, *s;
	int n, nfmt;

	nfmt = f-&gt;nfmt;
	for(;;){
		if(f-&gt;runes){
			rt = (Rune*)f-&gt;to;
			rs = (Rune*)f-&gt;stop;
			while((r = *(uchar*)fmt) &amp;&amp; r != &#39;%&#39;){
				if(r &lt; Runeself)
					fmt++;
				else{
					fmt += chartorune(&amp;rune, fmt);
					r = rune;
				}
				FMTRCHAR(f, rt, rs, r);
			}
			fmt++;
			f-&gt;nfmt += rt - (Rune *)f-&gt;to;
			f-&gt;to = rt;
			if(!r)
				return f-&gt;nfmt - nfmt;
			f-&gt;stop = rs;
		}else{
			t = (char*)f-&gt;to;
			s = (char*)f-&gt;stop;
			while((r = *(uchar*)fmt) &amp;&amp; r != &#39;%&#39;){
				if(r &lt; Runeself){
					FMTCHAR(f, t, s, r);
					fmt++;
				}else{
					n = chartorune(&amp;rune, fmt);
					if(t + n &gt; s){
						t = (char*)__fmtflush(f, t, n);
						if(t != nil)
							s = (char*)f-&gt;stop;
						else
							return -1;
					}
					while(n--)
						*t++ = *fmt++;
				}
			}
			fmt++;
			f-&gt;nfmt += t - (char *)f-&gt;to;
			f-&gt;to = t;
			if(!r)
				return f-&gt;nfmt - nfmt;
			f-&gt;stop = s;
		}

		fmt = (char*)__fmtdispatch(f, fmt, 0);
		if(fmt == nil)
			return -1;
	}
}

void *
__fmtflush(Fmt *f, void *t, int len)
{
	if(f-&gt;runes)
		f-&gt;nfmt += (Rune*)t - (Rune*)f-&gt;to;
	else
		f-&gt;nfmt += (char*)t - (char *)f-&gt;to;
	f-&gt;to = t;
	if(f-&gt;flush == 0 || (*f-&gt;flush)(f) == 0 || (char*)f-&gt;to + len &gt; (char*)f-&gt;stop){
		f-&gt;stop = f-&gt;to;
		return nil;
	}
	return f-&gt;to;
}

/*
 * put a formatted block of memory sz bytes long of n runes into the output buffer,
 * left/right justified in a field of at least f-&gt;width characters (if FmtWidth is set)
 */
int
__fmtpad(Fmt *f, int n)
{
	char *t, *s;
	int i;

	t = (char*)f-&gt;to;
	s = (char*)f-&gt;stop;
	for(i = 0; i &lt; n; i++)
		FMTCHAR(f, t, s, &#39; &#39;);
	f-&gt;nfmt += t - (char *)f-&gt;to;
	f-&gt;to = t;
	return 0;
}

int
__rfmtpad(Fmt *f, int n)
{
	Rune *t, *s;
	int i;

	t = (Rune*)f-&gt;to;
	s = (Rune*)f-&gt;stop;
	for(i = 0; i &lt; n; i++)
		FMTRCHAR(f, t, s, &#39; &#39;);
	f-&gt;nfmt += t - (Rune *)f-&gt;to;
	f-&gt;to = t;
	return 0;
}

int
__fmtcpy(Fmt *f, const void *vm, int n, int sz)
{
	Rune *rt, *rs, r;
	char *t, *s, *m, *me;
	ulong fl;
	int nc, w;

	m = (char*)vm;
	me = m + sz;
	fl = f-&gt;flags;
	w = 0;
	if(fl &amp; FmtWidth)
		w = f-&gt;width;
	if((fl &amp; FmtPrec) &amp;&amp; n &gt; f-&gt;prec)
		n = f-&gt;prec;
	if(f-&gt;runes){
		if(!(fl &amp; FmtLeft) &amp;&amp; __rfmtpad(f, w - n) &lt; 0)
			return -1;
		rt = (Rune*)f-&gt;to;
		rs = (Rune*)f-&gt;stop;
		for(nc = n; nc &gt; 0; nc--){
			r = *(uchar*)m;
			if(r &lt; Runeself)
				m++;
			else if((me - m) &gt;= UTFmax || fullrune(m, me-m))
				m += chartorune(&amp;r, m);
			else
				break;
			FMTRCHAR(f, rt, rs, r);
		}
		f-&gt;nfmt += rt - (Rune *)f-&gt;to;
		f-&gt;to = rt;
		if(fl &amp; FmtLeft &amp;&amp; __rfmtpad(f, w - n) &lt; 0)
			return -1;
	}else{
		if(!(fl &amp; FmtLeft) &amp;&amp; __fmtpad(f, w - n) &lt; 0)
			return -1;
		t = (char*)f-&gt;to;
		s = (char*)f-&gt;stop;
		for(nc = n; nc &gt; 0; nc--){
			r = *(uchar*)m;
			if(r &lt; Runeself)
				m++;
			else if((me - m) &gt;= UTFmax || fullrune(m, me-m))
				m += chartorune(&amp;r, m);
			else
				break;
			FMTRUNE(f, t, s, r);
		}
		f-&gt;nfmt += t - (char *)f-&gt;to;
		f-&gt;to = t;
		if(fl &amp; FmtLeft &amp;&amp; __fmtpad(f, w - n) &lt; 0)
			return -1;
	}
	return 0;
}

int
__fmtrcpy(Fmt *f, const void *vm, int n)
{
	Rune r, *m, *me, *rt, *rs;
	char *t, *s;
	ulong fl;
	int w;

	m = (Rune*)vm;
	fl = f-&gt;flags;
	w = 0;
	if(fl &amp; FmtWidth)
		w = f-&gt;width;
	if((fl &amp; FmtPrec) &amp;&amp; n &gt; f-&gt;prec)
		n = f-&gt;prec;
	if(f-&gt;runes){
		if(!(fl &amp; FmtLeft) &amp;&amp; __rfmtpad(f, w - n) &lt; 0)
			return -1;
		rt = (Rune*)f-&gt;to;
		rs = (Rune*)f-&gt;stop;
		for(me = m + n; m &lt; me; m++)
			FMTRCHAR(f, rt, rs, *m);
		f-&gt;nfmt += rt - (Rune *)f-&gt;to;
		f-&gt;to = rt;
		if(fl &amp; FmtLeft &amp;&amp; __rfmtpad(f, w - n) &lt; 0)
			return -1;
	}else{
		if(!(fl &amp; FmtLeft) &amp;&amp; __fmtpad(f, w - n) &lt; 0)
			return -1;
		t = (char*)f-&gt;to;
		s = (char*)f-&gt;stop;
		for(me = m + n; m &lt; me; m++){
			r = *m;
			FMTRUNE(f, t, s, r);
		}
		f-&gt;nfmt += t - (char *)f-&gt;to;
		f-&gt;to = t;
		if(fl &amp; FmtLeft &amp;&amp; __fmtpad(f, w - n) &lt; 0)
			return -1;
	}
	return 0;
}

/* fmt out one character */
int
__charfmt(Fmt *f)
{
	char x[1];

	x[0] = va_arg(f-&gt;args, int);
	f-&gt;prec = 1;
	return __fmtcpy(f, (const char*)x, 1, 1);
}

/* fmt out one rune */
int
__runefmt(Fmt *f)
{
	Rune x[1];

	x[0] = va_arg(f-&gt;args, int);
	return __fmtrcpy(f, (const void*)x, 1);
}

/* public helper routine: fmt out a null terminated string already in hand */
int
fmtstrcpy(Fmt *f, char *s)
{
	int i, j;

	if(!s)
		return __fmtcpy(f, &#34;&lt;nil&gt;&#34;, 5, 5);
	/* if precision is specified, make sure we don&#39;t wander off the end */
	if(f-&gt;flags &amp; FmtPrec){
#ifdef PLAN9PORT
		Rune r;
		i = 0;
		for(j=0; j&lt;f-&gt;prec &amp;&amp; s[i]; j++)
			i += chartorune(&amp;r, s+i);
#else
		/* ANSI requires precision in bytes, not Runes */
		for(i=0; i&lt;f-&gt;prec; i++)
			if(s[i] == 0)
				break;
		j = utfnlen(s, i);	/* won&#39;t print partial at end */
#endif
		return __fmtcpy(f, s, j, i);
	}
	return __fmtcpy(f, s, utflen(s), strlen(s));
}

/* fmt out a null terminated utf string */
int
__strfmt(Fmt *f)
{
	char *s;

	s = va_arg(f-&gt;args, char *);
	return fmtstrcpy(f, s);
}

/* public helper routine: fmt out a null terminated rune string already in hand */
int
fmtrunestrcpy(Fmt *f, Rune *s)
{
	Rune *e;
	int n, p;

	if(!s)
		return __fmtcpy(f, &#34;&lt;nil&gt;&#34;, 5, 5);
	/* if precision is specified, make sure we don&#39;t wander off the end */
	if(f-&gt;flags &amp; FmtPrec){
		p = f-&gt;prec;
		for(n = 0; n &lt; p; n++)
			if(s[n] == 0)
				break;
	}else{
		for(e = s; *e; e++)
			;
		n = e - s;
	}
	return __fmtrcpy(f, s, n);
}

/* fmt out a null terminated rune string */
int
__runesfmt(Fmt *f)
{
	Rune *s;

	s = va_arg(f-&gt;args, Rune *);
	return fmtrunestrcpy(f, s);
}

/* fmt a % */
int
__percentfmt(Fmt *f)
{
	Rune x[1];

	x[0] = f-&gt;r;
	f-&gt;prec = 1;
	return __fmtrcpy(f, (const void*)x, 1);
}

/* fmt an integer */
int
__ifmt(Fmt *f)
{
	char buf[140], *p, *conv;
	/* 140: for 64 bits of binary + 3-byte sep every 4 digits */
	uvlong vu;
	ulong u;
	int neg, base, i, n, fl, w, isv;
	int ndig, len, excess, bytelen;
	char *grouping;
	char *thousands;

	neg = 0;
	fl = f-&gt;flags;
	isv = 0;
	vu = 0;
	u = 0;
#ifndef PLAN9PORT
	/*
	 * Unsigned verbs for ANSI C
	 */
	switch(f-&gt;r){
	case &#39;o&#39;:
	case &#39;p&#39;:
	case &#39;u&#39;:
	case &#39;x&#39;:
	case &#39;X&#39;:
		fl |= FmtUnsigned;
		fl &amp;= ~(FmtSign|FmtSpace);
		break;
	}
#endif
	if(f-&gt;r == &#39;p&#39;){
		u = (ulong)va_arg(f-&gt;args, void*);
		f-&gt;r = &#39;x&#39;;
		fl |= FmtUnsigned;
	}else if(fl &amp; FmtVLong){
		isv = 1;
		if(fl &amp; FmtUnsigned)
			vu = va_arg(f-&gt;args, uvlong);
		else
			vu = va_arg(f-&gt;args, vlong);
	}else if(fl &amp; FmtLong){
		if(fl &amp; FmtUnsigned)
			u = va_arg(f-&gt;args, ulong);
		else
			u = va_arg(f-&gt;args, long);
	}else if(fl &amp; FmtByte){
		if(fl &amp; FmtUnsigned)
			u = (uchar)va_arg(f-&gt;args, int);
		else
			u = (char)va_arg(f-&gt;args, int);
	}else if(fl &amp; FmtShort){
		if(fl &amp; FmtUnsigned)
			u = (ushort)va_arg(f-&gt;args, int);
		else
			u = (short)va_arg(f-&gt;args, int);
	}else{
		if(fl &amp; FmtUnsigned)
			u = va_arg(f-&gt;args, uint);
		else
			u = va_arg(f-&gt;args, int);
	}
	conv = &#34;0123456789abcdef&#34;;
	grouping = &#34;\4&#34;;	/* for hex, octal etc. (undefined by spec but nice) */
	thousands = f-&gt;thousands;
	switch(f-&gt;r){
	case &#39;d&#39;:
	case &#39;i&#39;:
	case &#39;u&#39;:
		base = 10;
		grouping = f-&gt;grouping;
		break;
	case &#39;X&#39;:
		conv = &#34;0123456789ABCDEF&#34;;
		/* fall through */
	case &#39;x&#39;:
		base = 16;
		thousands = &#34;:&#34;;
		break;
	case &#39;b&#39;:
		base = 2;
		thousands = &#34;:&#34;;
		break;
	case &#39;o&#39;:
		base = 8;
		break;
	default:
		return -1;
	}
	if(!(fl &amp; FmtUnsigned)){
		if(isv &amp;&amp; (vlong)vu &lt; 0){
			vu = -(vlong)vu;
			neg = 1;
		}else if(!isv &amp;&amp; (long)u &lt; 0){
			u = -(long)u;
			neg = 1;
		}
	}
	p = buf + sizeof buf - 1;
	n = 0;	/* in runes */
	excess = 0;	/* number of bytes &gt; number runes */
	ndig = 0;
	len = utflen(thousands);
	bytelen = strlen(thousands);
	if(isv){
		while(vu){
			i = vu % base;
			vu /= base;
			if((fl &amp; FmtComma) &amp;&amp; n % 4 == 3){
				*p-- = &#39;,&#39;;
				n++;
			}
			if((fl &amp; FmtApost) &amp;&amp; __needsep(&amp;ndig, &amp;grouping)){
				n += len;
				excess += bytelen - len;
				p -= bytelen;
				memmove(p+1, thousands, bytelen);
			}
			*p-- = conv[i];
			n++;
		}
	}else{
		while(u){
			i = u % base;
			u /= base;
			if((fl &amp; FmtComma) &amp;&amp; n % 4 == 3){
				*p-- = &#39;,&#39;;
				n++;
			}
			if((fl &amp; FmtApost) &amp;&amp; __needsep(&amp;ndig, &amp;grouping)){
				n += len;
				excess += bytelen - len;
				p -= bytelen;
				memmove(p+1, thousands, bytelen);
			}
			*p-- = conv[i];
			n++;
		}
	}
	if(n == 0){
		/*
		 * &#34;The result of converting a zero value with
		 * a precision of zero is no characters.&#34;  - ANSI
		 *
		 * &#34;For o conversion, # increases the precision, if and only if
		 * necessary, to force the first digit of the result to be a zero
		 * (if the value and precision are both 0, a single 0 is printed).&#34; - ANSI
		 */
		if(!(fl &amp; FmtPrec) || f-&gt;prec != 0 || (f-&gt;r == &#39;o&#39; &amp;&amp; (fl &amp; FmtSharp))){
			*p-- = &#39;0&#39;;
			n = 1;
			if(fl &amp; FmtApost)
				__needsep(&amp;ndig, &amp;grouping);
		}

		/*
		 * Zero values don&#39;t get 0x.
		 */
		if(f-&gt;r == &#39;x&#39; || f-&gt;r == &#39;X&#39;)
			fl &amp;= ~FmtSharp;
	}
	for(w = f-&gt;prec; n &lt; w &amp;&amp; p &gt; buf+3; n++){
		if((fl &amp; FmtApost) &amp;&amp; __needsep(&amp;ndig, &amp;grouping)){
			n += len;
			excess += bytelen - len;
			p -= bytelen;
			memmove(p+1, thousands, bytelen);
		}
		*p-- = &#39;0&#39;;
	}
	if(neg || (fl &amp; (FmtSign|FmtSpace)))
		n++;
	if(fl &amp; FmtSharp){
		if(base == 16)
			n += 2;
		else if(base == 8){
			if(p[1] == &#39;0&#39;)
				fl &amp;= ~FmtSharp;
			else
				n++;
		}
	}
	if((fl &amp; FmtZero) &amp;&amp; !(fl &amp; (FmtLeft|FmtPrec))){
		w = 0;
		if(fl &amp; FmtWidth)
			w = f-&gt;width;
		for(; n &lt; w &amp;&amp; p &gt; buf+3; n++){
			if((fl &amp; FmtApost) &amp;&amp; __needsep(&amp;ndig, &amp;grouping)){
				n += len;
				excess += bytelen - len;
				p -= bytelen;
				memmove(p+1, thousands, bytelen);
			}
			*p-- = &#39;0&#39;;
		}
		f-&gt;flags &amp;= ~FmtWidth;
	}
	if(fl &amp; FmtSharp){
		if(base == 16)
			*p-- = f-&gt;r;
		if(base == 16 || base == 8)
			*p-- = &#39;0&#39;;
	}
	if(neg)
		*p-- = &#39;-&#39;;
	else if(fl &amp; FmtSign)
		*p-- = &#39;+&#39;;
	else if(fl &amp; FmtSpace)
		*p-- = &#39; &#39;;
	f-&gt;flags &amp;= ~FmtPrec;
	return __fmtcpy(f, p + 1, n, n + excess);
}

int
__countfmt(Fmt *f)
{
	void *p;
	ulong fl;

	fl = f-&gt;flags;
	p = va_arg(f-&gt;args, void*);
	if(fl &amp; FmtVLong){
		*(vlong*)p = f-&gt;nfmt;
	}else if(fl &amp; FmtLong){
		*(long*)p = f-&gt;nfmt;
	}else if(fl &amp; FmtByte){
		*(char*)p = f-&gt;nfmt;
	}else if(fl &amp; FmtShort){
		*(short*)p = f-&gt;nfmt;
	}else{
		*(int*)p = f-&gt;nfmt;
	}
	return 0;
}

int
__flagfmt(Fmt *f)
{
	switch(f-&gt;r){
	case &#39;,&#39;:
		f-&gt;flags |= FmtComma;
		break;
	case &#39;-&#39;:
		f-&gt;flags |= FmtLeft;
		break;
	case &#39;+&#39;:
		f-&gt;flags |= FmtSign;
		break;
	case &#39;#&#39;:
		f-&gt;flags |= FmtSharp;
		break;
	case &#39;\&#39;&#39;:
		f-&gt;flags |= FmtApost;
		break;
	case &#39; &#39;:
		f-&gt;flags |= FmtSpace;
		break;
	case &#39;u&#39;:
		f-&gt;flags |= FmtUnsigned;
		break;
	case &#39;h&#39;:
		if(f-&gt;flags &amp; FmtShort)
			f-&gt;flags |= FmtByte;
		f-&gt;flags |= FmtShort;
		break;
	case &#39;L&#39;:
		f-&gt;flags |= FmtLDouble;
		break;
	case &#39;l&#39;:
		if(f-&gt;flags &amp; FmtLong)
			f-&gt;flags |= FmtVLong;
		f-&gt;flags |= FmtLong;
		break;
	}
	return 1;
}

/* default error format */
int
__badfmt(Fmt *f)
{
	char x[2+UTFmax];
	int n;

	x[0] = &#39;%&#39;;
	n = 1 + runetochar(x+1, &amp;f-&gt;r);
	x[n++] = &#39;%&#39;;
	f-&gt;prec = n;
	__fmtcpy(f, (const void*)x, n, n);
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
