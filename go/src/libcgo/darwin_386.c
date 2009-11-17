<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/libcgo/darwin_386.c</title>

  <link rel="stylesheet" type="text/css" href="../../doc/style.css">
  <script type="text/javascript" src="../../doc/godocs.js"></script>

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
        <a href="../../index.html"><img src="../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../doc/install.html">Install Go</a></li>
    <li><a href="../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../pkg/index.html">Package documentation</a></li>
    <li><a href="../index.html">Source files</a></li>

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
  <h1 id="generatedHeader">Text file src/libcgo/darwin_386.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &lt;pthread.h&gt;
#include &#34;libcgo.h&#34;

static void* threadentry(void*);
static pthread_key_t k1, k2;

/* gccism: arrange for inittls to be called at dynamic load time */
static void inittls(void) __attribute__((constructor));

static void
inittls(void)
{
	uint32 x, y;
	pthread_key_t tofree[16], k;
	int i, ntofree;
	int havek1, havek2;

	/*
	 * Allocate thread-local storage slots for m, g.
	 * The key numbers start at 0x100, and we expect to be
	 * one of the early calls to pthread_key_create, so we
	 * should be able to get pretty low numbers.
	 *
	 * In Darwin/386 pthreads, %gs points at the thread
	 * structure, and each key is an index into the thread-local
	 * storage array that begins at offset 0x48 within in that structure.
	 * It may happen that we are not quite the first function to try
	 * to allocate thread-local storage keys, so instead of depending
	 * on getting 0x100 and 0x101, we try for 0x108 and 0x109,
	 * allocating keys until we get the ones we want and then freeing
	 * the ones we didn&#39;t want.
	 *
	 * Thus the final offsets to use in %gs references are
	 * 0x48+4*0x108 = 0x468 and 0x48+4*0x109 = 0x46c.
	 *
	 * The linker and runtime hard-code these constant offsets
	 * from %gs where we expect to find m and g.  The code
	 * below verifies that the constants are correct once it has
	 * obtained the keys.  Known to ../cmd/8l/obj.c:/468
	 * and to ../pkg/runtime/darwin/386/sys.s:/468
	 *
	 * This is truly disgusting and a bit fragile, but taking care
	 * of it here protects the rest of the system from damage.
	 * The alternative would be to use a global variable that
	 * held the offset and refer to that variable each time we
	 * need a %gs variable (m or g).  That approach would
	 * require an extra instruction and memory reference in
	 * every stack growth prolog and would also require
	 * rewriting the code that 8c generates for extern registers.
	 */
	havek1 = 0;
	havek2 = 0;
	ntofree = 0;
	while(!havek1 || !havek2) {
		if(pthread_key_create(&amp;k, nil) &lt; 0) {
			fprintf(stderr, &#34;libcgo: pthread_key_create failed\n&#34;);
			abort();
		}
		if(k == 0x108) {
			havek1 = 1;
			k1 = k;
			continue;
		}
		if(k == 0x109) {
			havek2 = 1;
			k2 = k;
			continue;
		}
		if(ntofree &gt;= nelem(tofree)) {
			fprintf(stderr, &#34;libcgo: could not obtain pthread_keys\n&#34;);
			fprintf(stderr, &#34;\twanted 0x108 and 0x109\n&#34;);
			fprintf(stderr, &#34;\tgot&#34;);
			for(i=0; i&lt;ntofree; i++)
				fprintf(stderr, &#34; %#x&#34;, tofree[i]);
			fprintf(stderr, &#34;\n&#34;);
			abort();
		}
		tofree[ntofree++] = k;
	}

	for(i=0; i&lt;ntofree; i++)
		pthread_key_delete(tofree[i]);

	/*
	 * We got the keys we wanted.  Make sure that we observe
	 * updates to k1 at 0x468, to verify that the TLS array
	 * offset from %gs hasn&#39;t changed.
	 */
	pthread_setspecific(k1, (void*)0x12345678);
	asm volatile(&#34;movl %%gs:0x468, %0&#34; : &#34;=r&#34;(x));

	pthread_setspecific(k1, (void*)0x87654321);
	asm volatile(&#34;movl %%gs:0x468, %0&#34; : &#34;=r&#34;(y));

	if(x != 0x12345678 || y != 0x87654321) {
		printf(&#34;libcgo: thread-local storage %#x not at %%gs:0x468 - x=%#x y=%#x\n&#34;, k1, x, y);
		abort();
	}
}

void
initcgo(void)
{
}

void
libcgo_sys_thread_start(ThreadStart *ts)
{
	pthread_attr_t attr;
	pthread_t p;
	size_t size;

	pthread_attr_init(&amp;attr);
	pthread_attr_getstacksize(&amp;attr, &amp;size);
	ts-&gt;g-&gt;stackguard = size;
	pthread_create(&amp;p, &amp;attr, threadentry, ts);
}

static void*
threadentry(void *v)
{
	ThreadStart ts;

	ts = *(ThreadStart*)v;
	free(v);

	ts.g-&gt;stackbase = (uintptr)&amp;ts;

	/*
	 * libcgo_sys_thread_start set stackguard to stack size;
	 * change to actual guard pointer.
	 */
	ts.g-&gt;stackguard = (uintptr)&amp;ts - ts.g-&gt;stackguard + 4096;

	pthread_setspecific(k1, (void*)ts.g);
	pthread_setspecific(k2, (void*)ts.m);

	crosscall_386(ts.fn);
	return nil;
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
