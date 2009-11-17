<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/lib9/fmt/fmtquote.c</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/lib9/fmt/fmtquote.c</h1>

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

/*
 * How many bytes of output UTF will be produced by quoting (if necessary) this string?
 * How many runes? How much of the input will be consumed?
 * The parameter q is filled in by __quotesetup.
 * The string may be UTF or Runes (s or r).
 * Return count does not include NUL.
 * Terminate the scan at the first of:
 *	NUL in input
 *	count exceeded in input
 *	count exceeded on output
 * *ninp is set to number of input bytes accepted.
 * nin may be &lt;0 initially, to avoid checking input by count.
 */
void
__quotesetup(char *s, Rune *r, int nin, int nout, Quoteinfo *q, int sharp, int runesout)
{
	int w;
	Rune c;

	q-&gt;quoted = 0;
	q-&gt;nbytesout = 0;
	q-&gt;nrunesout = 0;
	q-&gt;nbytesin = 0;
	q-&gt;nrunesin = 0;
	if(sharp || nin==0 || (s &amp;&amp; *s==&#39;\0&#39;) || (r &amp;&amp; *r==&#39;\0&#39;)){
		if(nout &lt; 2)
			return;
		q-&gt;quoted = 1;
		q-&gt;nbytesout = 2;
		q-&gt;nrunesout = 2;
	}
	for(; nin!=0; nin--){
		if(s)
			w = chartorune(&amp;c, s);
		else{
			c = *r;
			w = runelen(c);
		}

		if(c == &#39;\0&#39;)
			break;
		if(runesout){
			if(q-&gt;nrunesout+1 &gt; nout)
				break;
		}else{
			if(q-&gt;nbytesout+w &gt; nout)
				break;
		}

		if((c &lt;= L&#39; &#39;) || (c == L&#39;\&#39;&#39;) || (fmtdoquote!=nil &amp;&amp; fmtdoquote(c))){
			if(!q-&gt;quoted){
				if(runesout){
					if(1+q-&gt;nrunesout+1+1 &gt; nout)	/* no room for quotes */
						break;
				}else{
					if(1+q-&gt;nbytesout+w+1 &gt; nout)	/* no room for quotes */
						break;
				}
				q-&gt;nrunesout += 2;	/* include quotes */
				q-&gt;nbytesout += 2;	/* include quotes */
				q-&gt;quoted = 1;
			}
			if(c == &#39;\&#39;&#39;)	{
				if(runesout){
					if(1+q-&gt;nrunesout+1 &gt; nout)	/* no room for quotes */
						break;
				}else{
					if(1+q-&gt;nbytesout+w &gt; nout)	/* no room for quotes */
						break;
				}
				q-&gt;nbytesout++;
				q-&gt;nrunesout++;	/* quotes reproduce as two characters */
			}
		}

		/* advance input */
		if(s)
			s += w;
		else
			r++;
		q-&gt;nbytesin += w;
		q-&gt;nrunesin++;

		/* advance output */
		q-&gt;nbytesout += w;
		q-&gt;nrunesout++;

#ifndef PLAN9PORT
		/* ANSI requires precision in bytes, not Runes. */
		nin-= w-1;	/* and then n-- in the loop */
#endif
	}
}

static int
qstrfmt(char *sin, Rune *rin, Quoteinfo *q, Fmt *f)
{
	Rune r, *rm, *rme;
	char *t, *s, *m, *me;
	Rune *rt, *rs;
	ulong fl;
	int nc, w;

	m = sin;
	me = m + q-&gt;nbytesin;
	rm = rin;
	rme = rm + q-&gt;nrunesin;

	fl = f-&gt;flags;
	w = 0;
	if(fl &amp; FmtWidth)
		w = f-&gt;width;
	if(f-&gt;runes){
		if(!(fl &amp; FmtLeft) &amp;&amp; __rfmtpad(f, w - q-&gt;nrunesout) &lt; 0)
			return -1;
	}else{
		if(!(fl &amp; FmtLeft) &amp;&amp; __fmtpad(f, w - q-&gt;nbytesout) &lt; 0)
			return -1;
	}
	t = (char*)f-&gt;to;
	s = (char*)f-&gt;stop;
	rt = (Rune*)f-&gt;to;
	rs = (Rune*)f-&gt;stop;
	if(f-&gt;runes)
		FMTRCHAR(f, rt, rs, &#39;\&#39;&#39;);
	else
		FMTRUNE(f, t, s, &#39;\&#39;&#39;);
	for(nc = q-&gt;nrunesin; nc &gt; 0; nc--){
		if(sin){
			r = *(uchar*)m;
			if(r &lt; Runeself)
				m++;
			else if((me - m) &gt;= UTFmax || fullrune(m, me-m))
				m += chartorune(&amp;r, m);
			else
				break;
		}else{
			if(rm &gt;= rme)
				break;
			r = *(uchar*)rm++;
		}
		if(f-&gt;runes){
			FMTRCHAR(f, rt, rs, r);
			if(r == &#39;\&#39;&#39;)
				FMTRCHAR(f, rt, rs, r);
		}else{
			FMTRUNE(f, t, s, r);
			if(r == &#39;\&#39;&#39;)
				FMTRUNE(f, t, s, r);
		}
	}

	if(f-&gt;runes){
		FMTRCHAR(f, rt, rs, &#39;\&#39;&#39;);
		USED(rs);
		f-&gt;nfmt += rt - (Rune *)f-&gt;to;
		f-&gt;to = rt;
		if(fl &amp; FmtLeft &amp;&amp; __rfmtpad(f, w - q-&gt;nrunesout) &lt; 0)
			return -1;
	}else{
		FMTRUNE(f, t, s, &#39;\&#39;&#39;);
		USED(s);
		f-&gt;nfmt += t - (char *)f-&gt;to;
		f-&gt;to = t;
		if(fl &amp; FmtLeft &amp;&amp; __fmtpad(f, w - q-&gt;nbytesout) &lt; 0)
			return -1;
	}
	return 0;
}

int
__quotestrfmt(int runesin, Fmt *f)
{
	int nin, outlen;
	Rune *r;
	char *s;
	Quoteinfo q;

	nin = -1;
	if(f-&gt;flags&amp;FmtPrec)
		nin = f-&gt;prec;
	if(runesin){
		r = va_arg(f-&gt;args, Rune *);
		s = nil;
	}else{
		s = va_arg(f-&gt;args, char *);
		r = nil;
	}
	if(!s &amp;&amp; !r)
		return __fmtcpy(f, (void*)&#34;&lt;nil&gt;&#34;, 5, 5);

	if(f-&gt;flush)
		outlen = 0x7FFFFFFF;	/* if we can flush, no output limit */
	else if(f-&gt;runes)
		outlen = (Rune*)f-&gt;stop - (Rune*)f-&gt;to;
	else
		outlen = (char*)f-&gt;stop - (char*)f-&gt;to;

	__quotesetup(s, r, nin, outlen, &amp;q, f-&gt;flags&amp;FmtSharp, f-&gt;runes);
/*print(&#34;bytes in %d bytes out %d runes in %d runesout %d\n&#34;, q.nbytesin, q.nbytesout, q.nrunesin, q.nrunesout); */

	if(runesin){
		if(!q.quoted)
			return __fmtrcpy(f, r, q.nrunesin);
		return qstrfmt(nil, r, &amp;q, f);
	}

	if(!q.quoted)
		return __fmtcpy(f, s, q.nrunesin, q.nbytesin);
	return qstrfmt(s, nil, &amp;q, f);
}

int
quotestrfmt(Fmt *f)
{
	return __quotestrfmt(0, f);
}

int
quoterunestrfmt(Fmt *f)
{
	return __quotestrfmt(1, f);
}

void
quotefmtinstall(void)
{
	fmtinstall(&#39;q&#39;, quotestrfmt);
	fmtinstall(&#39;Q&#39;, quoterunestrfmt);
}

int
__needsquotes(char *s, int *quotelenp)
{
	Quoteinfo q;

	__quotesetup(s, nil, -1, 0x7FFFFFFF, &amp;q, 0, 0);
	*quotelenp = q.nbytesout;

	return q.quoted;
}

int
__runeneedsquotes(Rune *r, int *quotelenp)
{
	Quoteinfo q;

	__quotesetup(nil, r, -1, 0x7FFFFFFF, &amp;q, 0, 0);
	*quotelenp = q.nrunesout;

	return q.quoted;
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
