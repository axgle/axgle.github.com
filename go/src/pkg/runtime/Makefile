<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/Makefile</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/Makefile</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=runtime

# Set SIZE to 32 or 64.
SIZE_386=32
SIZE_amd64=64
SIZE_arm=32
SIZE=$(SIZE_$(GOARCH))

# Setup CFLAGS.  Add -D_64BIT on 64-bit platforms (sorry).
CFLAGS_64=-D_64BIT
# TODO(kaib): fix register allocation to honor extern register so we
# can enable optimizations again.
CFLAGS_arm=-N
CFLAGS=-I$(GOOS) -I$(GOOS)/$(GOARCH) -wF $(CFLAGS_$(SIZE)) $(CFLAGS_$(GOARCH))

GOFILES=\
	extern.go\
	type.go\

# 386-specific object files
OFILES_386=\
	vlop.$O\
	vlrt.$O\

GOARM ?= 6

# arm-specific object files
OFILES_arm=\
	cas$(GOARM).$O\
	memset.$O\
	vlop.$O\
	vlrt.$O\

OFILES=\
	asm.$O\
	cgocall.$O\
	chan.$O\
	closure.$O\
	float.$O\
	hashmap.$O\
	iface.$O\
	malloc.$O\
	mcache.$O\
	mcentral.$O\
	mem.$O\
	mfixalloc.$O\
	mgc0.$O\
	mheap.$O\
	mheapmap$(SIZE).$O\
	msize.$O\
	print.$O\
	proc.$O\
	reflect.$O\
	rune.$O\
	runtime.$O\
	rt0.$O\
	sema.$O\
	signal.$O\
	slice.$O\
	string.$O\
	symtab.$O\
	sys.$O\
	thread.$O\
	traceback.$O\
	$(OFILES_$(GOARCH))\

HFILES=\
	cgocall.h\
	runtime.h\
	hashmap.h\
	malloc.h\
	$(GOARCH)/asm.h\
	$(GOOS)/os.h\
	$(GOOS)/$(GOARCH)/defs.h\

include $(GOROOT)/src/Make.pkg

clean: clean-local

clean-local:
	rm -f cgo2c */asm.h runtime.acid.*

$(GOARCH)/asm.h: mkasmh.sh runtime.acid.$(GOARCH)
	./mkasmh.sh &gt;$@.x
	mv -f $@.x $@

cgo2c: cgo2c.c
	quietgcc -o $@ $&lt;

%.c:	%.cgo cgo2c
	./cgo2c $&lt; &gt; $@.tmp
	mv -f $@.tmp $@

%.$O:	$(GOARCH)/%.c
	$(CC) $(CFLAGS) $&lt;

%.$O:	$(GOOS)/%.c
	$(CC) $(CFLAGS) $&lt;

%.$O:	$(GOOS)/$(GOARCH)/%.c
	$(CC) $(CFLAGS) $&lt;

%.$O:	$(GOARCH)/%.s $(GOARCH)/asm.h
	$(AS) $&lt;

%.$O:	$(GOOS)/$(GOARCH)/%.s $(GOARCH)/asm.h
	$(AS) $&lt;

# for discovering offsets inside structs when debugging
runtime.acid.$(GOARCH): runtime.h proc.c
	$(CC) -a proc.c &gt;$@
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
