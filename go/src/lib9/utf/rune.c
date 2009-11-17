<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/lib9/utf/rune.c</title>

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
  <h1 id="generatedHeader">Text file src/lib9/utf/rune.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
/*
 * The authors of this software are Rob Pike and Ken Thompson.
 *              Copyright (c) 2002 by Lucent Technologies.
 *              Portions Copyright (c) 2009 The Go Authors.  All rights reserved.
 * Permission to use, copy, modify, and distribute this software for any
 * purpose without fee is hereby granted, provided that this entire notice
 * is included in all copies of any software which is or includes a copy
 * or modification of this software and in all copies of the supporting
 * documentation for such software.
 * THIS SOFTWARE IS BEING PROVIDED &#34;AS IS&#34;, WITHOUT ANY EXPRESS OR IMPLIED
 * WARRANTY.  IN PARTICULAR, NEITHER THE AUTHORS NOR LUCENT TECHNOLOGIES MAKE ANY
 * REPRESENTATION OR WARRANTY OF ANY KIND CONCERNING THE MERCHANTABILITY
 * OF THIS SOFTWARE OR ITS FITNESS FOR ANY PARTICULAR PURPOSE.
 */
#include &lt;stdarg.h&gt;
#include &lt;string.h&gt;
#include &#34;utf.h&#34;
#include &#34;utfdef.h&#34;

enum
{
	Bit1	= 7,
	Bitx	= 6,
	Bit2	= 5,
	Bit3	= 4,
	Bit4	= 3,
	Bit5	= 2,

	T1	= ((1&lt;&lt;(Bit1+1))-1) ^ 0xFF,	/* 0000 0000 */
	Tx	= ((1&lt;&lt;(Bitx+1))-1) ^ 0xFF,	/* 1000 0000 */
	T2	= ((1&lt;&lt;(Bit2+1))-1) ^ 0xFF,	/* 1100 0000 */
	T3	= ((1&lt;&lt;(Bit3+1))-1) ^ 0xFF,	/* 1110 0000 */
	T4	= ((1&lt;&lt;(Bit4+1))-1) ^ 0xFF,	/* 1111 0000 */
	T5	= ((1&lt;&lt;(Bit5+1))-1) ^ 0xFF,	/* 1111 1000 */

	Rune1	= (1&lt;&lt;(Bit1+0*Bitx))-1,		/* 0000 0000 0111 1111 */
	Rune2	= (1&lt;&lt;(Bit2+1*Bitx))-1,		/* 0000 0111 1111 1111 */
	Rune3	= (1&lt;&lt;(Bit3+2*Bitx))-1,		/* 1111 1111 1111 1111 */
	Rune4	= (1&lt;&lt;(Bit4+3*Bitx))-1,
                                        /* 0001 1111 1111 1111 1111 1111 */

	Maskx	= (1&lt;&lt;Bitx)-1,			/* 0011 1111 */
	Testx	= Maskx ^ 0xFF,			/* 1100 0000 */

	Bad	= Runeerror,
};

/*
 * Modified by Wei-Hwa Huang, Google Inc., on 2004-09-24
 * This is a slower but &#34;safe&#34; version of the old chartorune
 * that works on strings that are not necessarily null-terminated.
 *
 * If you know for sure that your string is null-terminated,
 * chartorune will be a bit faster.
 *
 * It is guaranteed not to attempt to access &#34;length&#34;
 * past the incoming pointer.  This is to avoid
 * possible access violations.  If the string appears to be
 * well-formed but incomplete (i.e., to get the whole Rune
 * we&#39;d need to read past str+length) then we&#39;ll set the Rune
 * to Bad and return 0.
 *
 * Note that if we have decoding problems for other
 * reasons, we return 1 instead of 0.
 */
int
charntorune(Rune *rune, const char *str, int length)
{
	int c, c1, c2, c3;
	long l;

	/* When we&#39;re not allowed to read anything */
	if(length &lt;= 0) {
		goto badlen;
	}

	/*
	 * one character sequence (7-bit value)
	 *	00000-0007F =&gt; T1
	 */
	c = *(uchar*)str;
	if(c &lt; Tx) {
		*rune = c;
		return 1;
	}

	// If we can&#39;t read more than one character we must stop
	if(length &lt;= 1) {
		goto badlen;
	}

	/*
	 * two character sequence (11-bit value)
	 *	0080-07FF =&gt; T2 Tx
	 */
	c1 = *(uchar*)(str+1) ^ Tx;
	if(c1 &amp; Testx)
		goto bad;
	if(c &lt; T3) {
		if(c &lt; T2)
			goto bad;
		l = ((c &lt;&lt; Bitx) | c1) &amp; Rune2;
		if(l &lt;= Rune1)
			goto bad;
		*rune = l;
		return 2;
	}

	// If we can&#39;t read more than two characters we must stop
	if(length &lt;= 2) {
		goto badlen;
	}

	/*
	 * three character sequence (16-bit value)
	 *	0800-FFFF =&gt; T3 Tx Tx
	 */
	c2 = *(uchar*)(str+2) ^ Tx;
	if(c2 &amp; Testx)
		goto bad;
	if(c &lt; T4) {
		l = ((((c &lt;&lt; Bitx) | c1) &lt;&lt; Bitx) | c2) &amp; Rune3;
		if(l &lt;= Rune2)
			goto bad;
		*rune = l;
		return 3;
	}

	if (length &lt;= 3)
		goto badlen;

	/*
	 * four character sequence (21-bit value)
	 *	10000-1FFFFF =&gt; T4 Tx Tx Tx
	 */
	c3 = *(uchar*)(str+3) ^ Tx;
	if (c3 &amp; Testx)
		goto bad;
	if (c &lt; T5) {
		l = ((((((c &lt;&lt; Bitx) | c1) &lt;&lt; Bitx) | c2) &lt;&lt; Bitx) | c3) &amp; Rune4;
		if (l &lt;= Rune3)
			goto bad;
		*rune = l;
		return 4;
	}

	// Support for 5-byte or longer UTF-8 would go here, but
	// since we don&#39;t have that, we&#39;ll just fall through to bad.

	/*
	 * bad decoding
	 */
bad:
	*rune = Bad;
	return 1;
badlen:
	*rune = Bad;
	return 0;

}


/*
 * This is the older &#34;unsafe&#34; version, which works fine on
 * null-terminated strings.
 */
int
chartorune(Rune *rune, const char *str)
{
	int c, c1, c2, c3;
	long l;

	/*
	 * one character sequence
	 *	00000-0007F =&gt; T1
	 */
	c = *(uchar*)str;
	if(c &lt; Tx) {
		*rune = c;
		return 1;
	}

	/*
	 * two character sequence
	 *	0080-07FF =&gt; T2 Tx
	 */
	c1 = *(uchar*)(str+1) ^ Tx;
	if(c1 &amp; Testx)
		goto bad;
	if(c &lt; T3) {
		if(c &lt; T2)
			goto bad;
		l = ((c &lt;&lt; Bitx) | c1) &amp; Rune2;
		if(l &lt;= Rune1)
			goto bad;
		*rune = l;
		return 2;
	}

	/*
	 * three character sequence
	 *	0800-FFFF =&gt; T3 Tx Tx
	 */
	c2 = *(uchar*)(str+2) ^ Tx;
	if(c2 &amp; Testx)
		goto bad;
	if(c &lt; T4) {
		l = ((((c &lt;&lt; Bitx) | c1) &lt;&lt; Bitx) | c2) &amp; Rune3;
		if(l &lt;= Rune2)
			goto bad;
		*rune = l;
		return 3;
	}

	/*
	 * four character sequence (21-bit value)
	 *	10000-1FFFFF =&gt; T4 Tx Tx Tx
	 */
	c3 = *(uchar*)(str+3) ^ Tx;
	if (c3 &amp; Testx)
		goto bad;
	if (c &lt; T5) {
		l = ((((((c &lt;&lt; Bitx) | c1) &lt;&lt; Bitx) | c2) &lt;&lt; Bitx) | c3) &amp; Rune4;
		if (l &lt;= Rune3)
			goto bad;
		*rune = l;
		return 4;
	}

	/*
	 * Support for 5-byte or longer UTF-8 would go here, but
	 * since we don&#39;t have that, we&#39;ll just fall through to bad.
	 */

	/*
	 * bad decoding
	 */
bad:
	*rune = Bad;
	return 1;
}

int
isvalidcharntorune(const char* str, int length, Rune* rune, int* consumed) {
	*consumed = charntorune(rune, str, length);
	return *rune != Runeerror || *consumed == 3;
}

int
runetochar(char *str, const Rune *rune)
{
	/* Runes are signed, so convert to unsigned for range check. */
	unsigned long c;

	/*
	 * one character sequence
	 *	00000-0007F =&gt; 00-7F
	 */
	c = *rune;
	if(c &lt;= Rune1) {
		str[0] = c;
		return 1;
	}

	/*
	 * two character sequence
	 *	0080-07FF =&gt; T2 Tx
	 */
	if(c &lt;= Rune2) {
		str[0] = T2 | (c &gt;&gt; 1*Bitx);
		str[1] = Tx | (c &amp; Maskx);
		return 2;
	}

	/*
	 * If the Rune is out of range, convert it to the error rune.
	 * Do this test here because the error rune encodes to three bytes.
	 * Doing it earlier would duplicate work, since an out of range
	 * Rune wouldn&#39;t have fit in one or two bytes.
	 */
	if (c &gt; Runemax)
		c = Runeerror;

	/*
	 * three character sequence
	 *	0800-FFFF =&gt; T3 Tx Tx
	 */
	if (c &lt;= Rune3) {
		str[0] = T3 |  (c &gt;&gt; 2*Bitx);
		str[1] = Tx | ((c &gt;&gt; 1*Bitx) &amp; Maskx);
		str[2] = Tx |  (c &amp; Maskx);
		return 3;
	}

	/*
	 * four character sequence (21-bit value)
	 *     10000-1FFFFF =&gt; T4 Tx Tx Tx
	 */
	str[0] = T4 | (c &gt;&gt; 3*Bitx);
	str[1] = Tx | ((c &gt;&gt; 2*Bitx) &amp; Maskx);
	str[2] = Tx | ((c &gt;&gt; 1*Bitx) &amp; Maskx);
	str[3] = Tx | (c &amp; Maskx);
	return 4;
}

int
runelen(Rune rune)
{
	char str[10];

	return runetochar(str, &amp;rune);
}

int
runenlen(const Rune *r, int nrune)
{
	int nb, c;

	nb = 0;
	while(nrune--) {
		c = *r++;
		if (c &lt;= Rune1)
			nb++;
		else if (c &lt;= Rune2)
			nb += 2;
		else if (c &lt;= Rune3)
			nb += 3;
		else /* assert(c &lt;= Rune4) */
			nb += 4;
	}
	return nb;
}

int
fullrune(const char *str, int n)
{
	if (n &gt; 0) {
		int c = *(uchar*)str;
		if (c &lt; Tx)
			return 1;
		if (n &gt; 1) {
			if (c &lt; T3)
				return 1;
			if (n &gt; 2) {
				if (c &lt; T4 || n &gt; 3)
					return 1;
			}
		}
	}
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
