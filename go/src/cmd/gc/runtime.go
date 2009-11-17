<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/gc/runtime.go</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/gc/runtime.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package PACKAGE

<a id="L7"></a><span class="comment">// emitted by compiler, not referred to by go programs</span>

<a id="L9"></a>func mal(int32) *any
<a id="L10"></a>func throwindex()
<a id="L11"></a>func throwreturn()
<a id="L12"></a>func throwinit()
<a id="L13"></a>func panicl()

<a id="L15"></a>func printbool(bool)
<a id="L16"></a>func printfloat(float64)
<a id="L17"></a>func printint(int64)
<a id="L18"></a>func printuint(uint64)
<a id="L19"></a>func printstring(string)
<a id="L20"></a>func printpointer(any)
<a id="L21"></a>func printiface(any)
<a id="L22"></a>func printeface(any)
<a id="L23"></a>func printslice(any)
<a id="L24"></a>func printnl()
<a id="L25"></a>func printsp()

<a id="L27"></a>func catstring(string, string) string
<a id="L28"></a>func cmpstring(string, string) int
<a id="L29"></a>func slicestring(string, int, int) string
<a id="L30"></a>func indexstring(string, int) byte
<a id="L31"></a>func intstring(int64) string
<a id="L32"></a>func slicebytetostring([]byte) string
<a id="L33"></a>func sliceinttostring([]int) string
<a id="L34"></a>func stringiter(string, int) int
<a id="L35"></a>func stringiter2(string, int) (retk int, retv int)

<a id="L37"></a>func ifaceI2E(iface any) (ret any)
<a id="L38"></a>func ifaceE2I(typ *byte, iface any) (ret any)
<a id="L39"></a>func ifaceT2E(typ *byte, elem any) (ret any)
<a id="L40"></a>func ifaceE2T(typ *byte, elem any) (ret any)
<a id="L41"></a>func ifaceE2I2(typ *byte, iface any) (ret any, ok bool)
<a id="L42"></a>func ifaceE2T2(typ *byte, elem any) (ret any, ok bool)
<a id="L43"></a>func ifaceT2I(typ1 *byte, typ2 *byte, elem any) (ret any)
<a id="L44"></a>func ifaceI2T(typ *byte, iface any) (ret any)
<a id="L45"></a>func ifaceI2T2(typ *byte, iface any) (ret any, ok bool)
<a id="L46"></a>func ifaceI2I(typ *byte, iface any) (ret any)
<a id="L47"></a>func ifaceI2Ix(typ *byte, iface any) (ret any)
<a id="L48"></a>func ifaceI2I2(typ *byte, iface any) (ret any, ok bool)
<a id="L49"></a>func ifaceeq(i1 any, i2 any) (ret bool)
<a id="L50"></a>func efaceeq(i1 any, i2 any) (ret bool)
<a id="L51"></a>func ifacethash(i1 any) (ret uint32)
<a id="L52"></a>func efacethash(i1 any) (ret uint32)

<a id="L54"></a><span class="comment">// *byte is really *runtime.Type</span>
<a id="L55"></a>func makemap(key, val *byte, hint int) (hmap map[any]any)
<a id="L56"></a>func mapaccess1(hmap map[any]any, key any) (val any)
<a id="L57"></a>func mapaccess2(hmap map[any]any, key any) (val any, pres bool)
<a id="L58"></a>func mapassign1(hmap map[any]any, key any, val any)
<a id="L59"></a>func mapassign2(hmap map[any]any, key any, val any, pres bool)
<a id="L60"></a>func mapiterinit(hmap map[any]any, hiter *any)
<a id="L61"></a>func mapiternext(hiter *any)
<a id="L62"></a>func mapiter1(hiter *any) (key any)
<a id="L63"></a>func mapiter2(hiter *any) (key any, val any)

<a id="L65"></a><span class="comment">// *byte is really *runtime.Type</span>
<a id="L66"></a>func makechan(elem *byte, hint int) (hchan chan any)
<a id="L67"></a>func chanrecv1(hchan &lt;-chan any) (elem any)
<a id="L68"></a>func chanrecv2(hchan &lt;-chan any) (elem any, pres bool)
<a id="L69"></a>func chansend1(hchan chan&lt;- any, elem any)
<a id="L70"></a>func chansend2(hchan chan&lt;- any, elem any) (pres bool)
<a id="L71"></a>func closechan(hchan any)
<a id="L72"></a>func closedchan(hchan any) bool

<a id="L74"></a>func newselect(size int) (sel *byte)
<a id="L75"></a>func selectsend(sel *byte, hchan chan&lt;- any, elem any) (selected bool)
<a id="L76"></a>func selectrecv(sel *byte, hchan &lt;-chan any, elem *any) (selected bool)
<a id="L77"></a>func selectdefault(sel *byte) (selected bool)
<a id="L78"></a>func selectgo(sel *byte)

<a id="L80"></a>func makeslice(nel int, cap int, width int) (ary []any)
<a id="L81"></a>func sliceslice(old []any, lb int, hb int, width int) (ary []any)
<a id="L82"></a>func slicearray(old *any, nel int, lb int, hb int, width int) (ary []any)
<a id="L83"></a>func arraytoslice(old *any, nel int) (ary []any)

<a id="L85"></a>func closure() <span class="comment">// has args, but compiler fills in</span>

<a id="L87"></a><span class="comment">// only used on 32-bit</span>
<a id="L88"></a>func int64div(int64, int64) int64
<a id="L89"></a>func uint64div(uint64, uint64) uint64
<a id="L90"></a>func int64mod(int64, int64) int64
<a id="L91"></a>func uint64mod(uint64, uint64) uint64
<a id="L92"></a>func float64toint64(float64) int64
<a id="L93"></a>func int64tofloat64(int64) float64
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
