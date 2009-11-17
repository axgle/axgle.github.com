<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/runtime.c</title>

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
  <h1 id="generatedHeader">Text file src/pkg/runtime/runtime.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;runtime.h&#34;

int32	panicking	= 0;
int32	maxround	= sizeof(uintptr);
int32	fd		= 1;

int32
gotraceback(void)
{
	byte *p;

	p = getenv(&#34;GOTRACEBACK&#34;);
	if(p == nil || p[0] == &#39;\0&#39;)
		return 1;	// default is on
	return atoi(p);
}

void
runtime·panicl(int32 lno)
{
	uint8 *sp;

	fd = 2;
	if(panicking) {
		printf(&#34;double panic\n&#34;);
		exit(3);
	}
	panicking++;

	printf(&#34;\npanic PC=%X\n&#34;, (uint64)(uintptr)&amp;lno);
	sp = (uint8*)&amp;lno;
	if(gotraceback()){
		traceback(runtime·getcallerpc(&amp;lno), sp, g);
		tracebackothers(g);
	}
	breakpoint();  // so we can grab it in a debugger
	exit(2);
}

void
runtime·throwindex(void)
{
	throw(&#34;index out of range&#34;);
}

void
runtime·throwslice(void)
{
	throw(&#34;slice out of range&#34;);
}

void
runtime·throwreturn(void)
{
	throw(&#34;no return at end of a typed function&#34;);
}

void
runtime·throwinit(void)
{
	throw(&#34;recursive call during initialization&#34;);
}

void
throw(int8 *s)
{
	fd = 2;
	printf(&#34;throw: %s\n&#34;, s);
	runtime·panicl(-1);
	*(int32*)0 = 0;	// not reached
	exit(1);	// even more not reached
}

void
mcpy(byte *t, byte *f, uint32 n)
{
	while(n &gt; 0) {
		*t = *f;
		t++;
		f++;
		n--;
	}
}

int32
mcmp(byte *s1, byte *s2, uint32 n)
{
	uint32 i;
	byte c1, c2;

	for(i=0; i&lt;n; i++) {
		c1 = s1[i];
		c2 = s2[i];
		if(c1 &lt; c2)
			return -1;
		if(c1 &gt; c2)
			return +1;
	}
	return 0;
}


void
mmov(byte *t, byte *f, uint32 n)
{
	if(t &lt; f) {
		while(n &gt; 0) {
			*t = *f;
			t++;
			f++;
			n--;
		}
	} else {
		t += n;
		f += n;
		while(n &gt; 0) {
			t--;
			f--;
			*t = *f;
			n--;
		}
	}
}

byte*
mchr(byte *p, byte c, byte *ep)
{
	for(; p &lt; ep; p++)
		if(*p == c)
			return p;
	return nil;
}

uint32
rnd(uint32 n, uint32 m)
{
	uint32 r;

	if(m &gt; maxround)
		m = maxround;
	r = n % m;
	if(r)
		n += m-r;
	return n;
}

static int32	argc;
static uint8**	argv;

Slice os·Args;
Slice os·Envs;

void
args(int32 c, uint8 **v)
{
	argc = c;
	argv = v;
}

void
goargs(void)
{
	String *gargv;
	String *genvv;
	int32 i, envc;

	for(envc=0; argv[argc+1+envc] != 0; envc++)
		;

	gargv = malloc(argc*sizeof gargv[0]);
	genvv = malloc(envc*sizeof genvv[0]);

	for(i=0; i&lt;argc; i++)
		gargv[i] = gostring(argv[i]);
	os·Args.array = (byte*)gargv;
	os·Args.len = argc;
	os·Args.cap = argc;

	for(i=0; i&lt;envc; i++)
		genvv[i] = gostring(argv[argc+1+i]);
	os·Envs.array = (byte*)genvv;
	os·Envs.len = envc;
	os·Envs.cap = envc;
}

byte*
getenv(int8 *s)
{
	int32 i, j, len;
	byte *v, *bs;
	String* envv;
	int32 envc;

	bs = (byte*)s;
	len = findnull(bs);
	envv = (String*)os·Envs.array;
	envc = os·Envs.len;
	for(i=0; i&lt;envc; i++){
		if(envv[i].len &lt;= len)
			continue;
		v = envv[i].str;
		for(j=0; j&lt;len; j++)
			if(bs[j] != v[j])
				goto nomatch;
		if(v[len] != &#39;=&#39;)
			goto nomatch;
		return v+len+1;
	nomatch:;
	}
	return nil;
}


int32
atoi(byte *p)
{
	int32 n;

	n = 0;
	while(&#39;0&#39; &lt;= *p &amp;&amp; *p &lt;= &#39;9&#39;)
		n = n*10 + *p++ - &#39;0&#39;;
	return n;
}

void
check(void)
{
	int8 a;
	uint8 b;
	int16 c;
	uint16 d;
	int32 e;
	uint32 f;
	int64 g;
	uint64 h;
	float32 i;
	float64 j;
	void* k;
	uint16* l;

	if(sizeof(a) != 1) throw(&#34;bad a&#34;);
	if(sizeof(b) != 1) throw(&#34;bad b&#34;);
	if(sizeof(c) != 2) throw(&#34;bad c&#34;);
	if(sizeof(d) != 2) throw(&#34;bad d&#34;);
	if(sizeof(e) != 4) throw(&#34;bad e&#34;);
	if(sizeof(f) != 4) throw(&#34;bad f&#34;);
	if(sizeof(g) != 8) throw(&#34;bad g&#34;);
	if(sizeof(h) != 8) throw(&#34;bad h&#34;);
	if(sizeof(i) != 4) throw(&#34;bad i&#34;);
	if(sizeof(j) != 8) throw(&#34;bad j&#34;);
	if(sizeof(k) != sizeof(uintptr)) throw(&#34;bad k&#34;);
	if(sizeof(l) != sizeof(uintptr)) throw(&#34;bad l&#34;);
//	prints(1&#34;check ok\n&#34;);

	uint32 z;
	z = 1;
	if(!cas(&amp;z, 1, 2))
		throw(&#34;cas1&#34;);
	if(z != 2)
		throw(&#34;cas2&#34;);

	z = 4;
	if(cas(&amp;z, 5, 6))
		throw(&#34;cas3&#34;);
	if(z != 4)
		throw(&#34;cas4&#34;);

	initsig();
}

/*
 * map and chan helpers for
 * dealing with unknown types
 */
static uintptr
memhash(uint32 s, void *a)
{
	byte *b;
	uintptr hash;

	b = a;
	if(sizeof(hash) == 4)
		hash = 2860486313U;
	else
		hash = 33054211828000289ULL;
	while(s &gt; 0) {
		if(sizeof(hash) == 4)
			hash = (hash ^ *b) * 3267000013UL;
		else
			hash = (hash ^ *b) * 23344194077549503ULL;
		b++;
		s--;
	}
	return hash;
}

static uint32
memequal(uint32 s, void *a, void *b)
{
	byte *ba, *bb;
	uint32 i;

	ba = a;
	bb = b;
	for(i=0; i&lt;s; i++)
		if(ba[i] != bb[i])
			return 0;
	return 1;
}

static void
memprint(uint32 s, void *a)
{
	uint64 v;

	v = 0xbadb00b;
	switch(s) {
	case 1:
		v = *(uint8*)a;
		break;
	case 2:
		v = *(uint16*)a;
		break;
	case 4:
		v = *(uint32*)a;
		break;
	case 8:
		v = *(uint64*)a;
		break;
	}
	runtime·printint(v);
}

static void
memcopy(uint32 s, void *a, void *b)
{
	byte *ba, *bb;
	uint32 i;

	ba = a;
	bb = b;
	if(bb == nil) {
		for(i=0; i&lt;s; i++)
			ba[i] = 0;
		return;
	}
	for(i=0; i&lt;s; i++)
		ba[i] = bb[i];
}

static uintptr
strhash(uint32 s, String *a)
{
	USED(s);
	return memhash((*a).len, (*a).str);
}

static uint32
strequal(uint32 s, String *a, String *b)
{
	USED(s);
	return cmpstring(*a, *b) == 0;
}

static void
strprint(uint32 s, String *a)
{
	USED(s);
	runtime·printstring(*a);
}

static uintptr
interhash(uint32 s, Iface *a)
{
	USED(s);
	return ifacehash(*a);
}

static void
interprint(uint32 s, Iface *a)
{
	USED(s);
	runtime·printiface(*a);
}

static uint32
interequal(uint32 s, Iface *a, Iface *b)
{
	USED(s);
	return ifaceeq(*a, *b);
}

static uintptr
nilinterhash(uint32 s, Eface *a)
{
	USED(s);
	return efacehash(*a);
}

static void
nilinterprint(uint32 s, Eface *a)
{
	USED(s);
	runtime·printeface(*a);
}

static uint32
nilinterequal(uint32 s, Eface *a, Eface *b)
{
	USED(s);
	return efaceeq(*a, *b);
}

uintptr
nohash(uint32 s, void *a)
{
	USED(s);
	USED(a);
	throw(&#34;hash of unhashable type&#34;);
	return 0;
}

uint32
noequal(uint32 s, void *a, void *b)
{
	USED(s);
	USED(a);
	USED(b);
	throw(&#34;comparing uncomparable types&#34;);
	return 0;
}

static void
noprint(uint32 s, void *a)
{
	USED(s);
	USED(a);
	throw(&#34;print of unprintable type&#34;);
}

static void
nocopy(uint32 s, void *a, void *b)
{
	USED(s);
	USED(a);
	USED(b);
	throw(&#34;copy of uncopyable type&#34;);
}

Alg
algarray[] =
{
[AMEM]	{ memhash, memequal, memprint, memcopy },
[ANOEQ]	{ nohash, noequal, memprint, memcopy },
[ASTRING]	{ strhash, strequal, strprint, memcopy },
[AINTER]		{ interhash, interequal, interprint, memcopy },
[ANILINTER]	{ nilinterhash, nilinterequal, nilinterprint, memcopy },
[AFAKE]	{ nohash, noequal, noprint, nocopy },
};

#pragma textflag 7
void
FLUSH(void *v)
{
	USED(v);
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
