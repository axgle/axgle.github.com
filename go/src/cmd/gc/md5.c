<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/md5.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/md5.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 64-bit MD5 (does full MD5 but returns 64 bits only).
// Translation of ../../pkg/crypto/md5/md5*.go.

#include &#34;go.h&#34;
#include &#34;md5.h&#34;

static int md5block(MD5 *dig, uchar *p, int nn);

enum {
	_Chunk = 64
};

#define _Init0 0x67452301
#define _Init1 0xEFCDAB89
#define _Init2 0x98BADCFE
#define _Init3 0x10325476

void
md5reset(MD5 *d)
{
	d-&gt;s[0] = _Init0;
	d-&gt;s[1] = _Init1;
	d-&gt;s[2] = _Init2;
	d-&gt;s[3] = _Init3;
	d-&gt;nx = 0;
	d-&gt;len = 0;
}

void
md5write(MD5 *d, uchar *p, int nn)
{
	int i, n;

	d-&gt;len += nn;
	if(d-&gt;nx &gt; 0) {
		n = nn;
		if(n &gt; _Chunk - d-&gt;nx)
			n = _Chunk - d-&gt;nx;
		for(i=0; i&lt;n; i++)
			d-&gt;x[d-&gt;nx+i] = p[i];
		d-&gt;nx += n;
		if(d-&gt;nx == _Chunk) {
			md5block(d, d-&gt;x, _Chunk);
			d-&gt;nx = 0;
		}
		p += n;
		nn -= n;
	}
	n = md5block(d, p, nn);
	p += n;
	nn -= n;
	if(nn &gt; 0) {
		for(i=0; i&lt;nn; i++)
			d-&gt;x[i] = p[i];
		d-&gt;nx = nn;
	}
}

uint64
md5sum(MD5 *d)
{
	uchar tmp[64];
	int i;
	uint64 len;

	// Padding.  Add a 1 bit and 0 bits until 56 bytes mod 64.
	len = d-&gt;len;
	memset(tmp, 0, sizeof tmp);
	tmp[0] = 0x80;
	if(len%64 &lt; 56)
		md5write(d, tmp, 56-len%64);
	else
		md5write(d, tmp, 64+56-len%64);

	// Length in bits.
	len &lt;&lt;= 3;
	for(i=0; i&lt;8; i++)
		tmp[i] = len&gt;&gt;(8*i);
	md5write(d, tmp, 8);

	if(d-&gt;nx != 0)
		fatal(&#34;md5sum&#34;);

	return d-&gt;s[0] | ((uint64)d-&gt;s[1]&lt;&lt;32);
}


// MD5 block step.
// In its own file so that a faster assembly or C version
// can be substituted easily.

// table[i] = int((1&lt;&lt;32) * abs(sin(i+1 radians))).
static uint32 table[64] = {
	// round 1
	0xd76aa478,
	0xe8c7b756,
	0x242070db,
	0xc1bdceee,
	0xf57c0faf,
	0x4787c62a,
	0xa8304613,
	0xfd469501,
	0x698098d8,
	0x8b44f7af,
	0xffff5bb1,
	0x895cd7be,
	0x6b901122,
	0xfd987193,
	0xa679438e,
	0x49b40821,

	// round 2
	0xf61e2562,
	0xc040b340,
	0x265e5a51,
	0xe9b6c7aa,
	0xd62f105d,
	0x2441453,
	0xd8a1e681,
	0xe7d3fbc8,
	0x21e1cde6,
	0xc33707d6,
	0xf4d50d87,
	0x455a14ed,
	0xa9e3e905,
	0xfcefa3f8,
	0x676f02d9,
	0x8d2a4c8a,

	// round3
	0xfffa3942,
	0x8771f681,
	0x6d9d6122,
	0xfde5380c,
	0xa4beea44,
	0x4bdecfa9,
	0xf6bb4b60,
	0xbebfbc70,
	0x289b7ec6,
	0xeaa127fa,
	0xd4ef3085,
	0x4881d05,
	0xd9d4d039,
	0xe6db99e5,
	0x1fa27cf8,
	0xc4ac5665,

	// round 4
	0xf4292244,
	0x432aff97,
	0xab9423a7,
	0xfc93a039,
	0x655b59c3,
	0x8f0ccc92,
	0xffeff47d,
	0x85845dd1,
	0x6fa87e4f,
	0xfe2ce6e0,
	0xa3014314,
	0x4e0811a1,
	0xf7537e82,
	0xbd3af235,
	0x2ad7d2bb,
	0xeb86d391,
};

static uint32 shift1[] = { 7, 12, 17, 22 };
static uint32 shift2[] = { 5, 9, 14, 20 };
static uint32 shift3[] = { 4, 11, 16, 23 };
static uint32 shift4[] = { 6, 10, 15, 21 };

static int
md5block(MD5 *dig, uchar *p, int nn)
{
	uint32 a, b, c, d, aa, bb, cc, dd;
	int i, j, n;
	uint32 X[16];

	a = dig-&gt;s[0];
	b = dig-&gt;s[1];
	c = dig-&gt;s[2];
	d = dig-&gt;s[3];
	n = 0;

	while(nn &gt;= _Chunk) {
		aa = a;
		bb = b;
		cc = c;
		dd = d;

		for(i=0; i&lt;16; i++) {
			j = i*4;
			X[i] = p[j] | (p[j+1]&lt;&lt;8) | (p[j+2]&lt;&lt;16) | (p[j+3]&lt;&lt;24);
		}

		// Round 1.
		for(i=0; i&lt;16; i++) {
			uint32 x, t, s, f;
			x = i;
			t = i;
			s = shift1[i%4];
			f = ((c ^ d) &amp; b) ^ d;
			a += f + X[x] + table[t];
			a = a&lt;&lt;s | a&gt;&gt;(32-s);
			a += b;

			t = d;
			d = c;
			c = b;
			b = a;
			a = t;
		}

		// Round 2.
		for(i=0; i&lt;16; i++) {
			uint32 x, t, s, g;

			x = (1+5*i)%16;
			t = 16+i;
			s = shift2[i%4];
			g = ((b ^ c) &amp; d) ^ c;
			a += g + X[x] + table[t];
			a = a&lt;&lt;s | a&gt;&gt;(32-s);
			a += b;

			t = d;
			d = c;
			c = b;
			b = a;
			a = t;
		}

		// Round 3.
		for(i=0; i&lt;16; i++) {
			uint32 x, t, s, h;

			x = (5+3*i)%16;
			t = 32+i;
			s = shift3[i%4];
			h = b ^ c ^ d;
			a += h + X[x] + table[t];
			a = a&lt;&lt;s | a&gt;&gt;(32-s);
			a += b;

			t = d;
			d = c;
			c = b;
			b = a;
			a = t;
		}

		// Round 4.
		for(i=0; i&lt;16; i++) {
			uint32 x, s, t, ii;

			x = (7*i)%16;
			s = shift4[i%4];
			t = 48+i;
			ii = c ^ (b | ~d);
			a += ii + X[x] + table[t];
			a = a&lt;&lt;s | a&gt;&gt;(32-s);
			a += b;

			t = d;
			d = c;
			c = b;
			b = a;
			a = t;
		}

		a += aa;
		b += bb;
		c += cc;
		d += dd;

		p += _Chunk;
		n += _Chunk;
		nn -= _Chunk;
	}

	dig-&gt;s[0] = a;
	dig-&gt;s[1] = b;
	dig-&gt;s[2] = c;
	dig-&gt;s[3] = d;
	return n;
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
