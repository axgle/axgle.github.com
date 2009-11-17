<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/builtin.c.boot</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/builtin.c.boot</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
char *runtimeimport =
	&#34;package runtime\n&#34;
	&#34;func runtime.mal (? int32) (? *any)\n&#34;
	&#34;func runtime.throwindex ()\n&#34;
	&#34;func runtime.throwreturn ()\n&#34;
	&#34;func runtime.throwinit ()\n&#34;
	&#34;func runtime.panicl ()\n&#34;
	&#34;func runtime.printbool (? bool)\n&#34;
	&#34;func runtime.printfloat (? float64)\n&#34;
	&#34;func runtime.printint (? int64)\n&#34;
	&#34;func runtime.printuint (? uint64)\n&#34;
	&#34;func runtime.printstring (? string)\n&#34;
	&#34;func runtime.printpointer (? any)\n&#34;
	&#34;func runtime.printiface (? any)\n&#34;
	&#34;func runtime.printeface (? any)\n&#34;
	&#34;func runtime.printslice (? any)\n&#34;
	&#34;func runtime.printnl ()\n&#34;
	&#34;func runtime.printsp ()\n&#34;
	&#34;func runtime.catstring (? string, ? string) (? string)\n&#34;
	&#34;func runtime.cmpstring (? string, ? string) (? int)\n&#34;
	&#34;func runtime.slicestring (? string, ? int, ? int) (? string)\n&#34;
	&#34;func runtime.indexstring (? string, ? int) (? uint8)\n&#34;
	&#34;func runtime.intstring (? int64) (? string)\n&#34;
	&#34;func runtime.slicebytetostring (? []uint8) (? string)\n&#34;
	&#34;func runtime.sliceinttostring (? []int) (? string)\n&#34;
	&#34;func runtime.stringiter (? string, ? int) (? int)\n&#34;
	&#34;func runtime.stringiter2 (? string, ? int) (retk int, retv int)\n&#34;
	&#34;func runtime.ifaceI2E (iface any) (ret any)\n&#34;
	&#34;func runtime.ifaceE2I (typ *uint8, iface any) (ret any)\n&#34;
	&#34;func runtime.ifaceT2E (typ *uint8, elem any) (ret any)\n&#34;
	&#34;func runtime.ifaceE2T (typ *uint8, elem any) (ret any)\n&#34;
	&#34;func runtime.ifaceE2I2 (typ *uint8, iface any) (ret any, ok bool)\n&#34;
	&#34;func runtime.ifaceE2T2 (typ *uint8, elem any) (ret any, ok bool)\n&#34;
	&#34;func runtime.ifaceT2I (typ1 *uint8, typ2 *uint8, elem any) (ret any)\n&#34;
	&#34;func runtime.ifaceI2T (typ *uint8, iface any) (ret any)\n&#34;
	&#34;func runtime.ifaceI2T2 (typ *uint8, iface any) (ret any, ok bool)\n&#34;
	&#34;func runtime.ifaceI2I (typ *uint8, iface any) (ret any)\n&#34;
	&#34;func runtime.ifaceI2Ix (typ *uint8, iface any) (ret any)\n&#34;
	&#34;func runtime.ifaceI2I2 (typ *uint8, iface any) (ret any, ok bool)\n&#34;
	&#34;func runtime.ifaceeq (i1 any, i2 any) (ret bool)\n&#34;
	&#34;func runtime.efaceeq (i1 any, i2 any) (ret bool)\n&#34;
	&#34;func runtime.ifacethash (i1 any) (ret uint32)\n&#34;
	&#34;func runtime.efacethash (i1 any) (ret uint32)\n&#34;
	&#34;func runtime.makemap (key *uint8, val *uint8, hint int) (hmap map[any] any)\n&#34;
	&#34;func runtime.mapaccess1 (hmap map[any] any, key any) (val any)\n&#34;
	&#34;func runtime.mapaccess2 (hmap map[any] any, key any) (val any, pres bool)\n&#34;
	&#34;func runtime.mapassign1 (hmap map[any] any, key any, val any)\n&#34;
	&#34;func runtime.mapassign2 (hmap map[any] any, key any, val any, pres bool)\n&#34;
	&#34;func runtime.mapiterinit (hmap map[any] any, hiter *any)\n&#34;
	&#34;func runtime.mapiternext (hiter *any)\n&#34;
	&#34;func runtime.mapiter1 (hiter *any) (key any)\n&#34;
	&#34;func runtime.mapiter2 (hiter *any) (key any, val any)\n&#34;
	&#34;func runtime.makechan (elem *uint8, hint int) (hchan chan any)\n&#34;
	&#34;func runtime.chanrecv1 (hchan &lt;-chan any) (elem any)\n&#34;
	&#34;func runtime.chanrecv2 (hchan &lt;-chan any) (elem any, pres bool)\n&#34;
	&#34;func runtime.chansend1 (hchan chan&lt;- any, elem any)\n&#34;
	&#34;func runtime.chansend2 (hchan chan&lt;- any, elem any) (pres bool)\n&#34;
	&#34;func runtime.closechan (hchan any)\n&#34;
	&#34;func runtime.closedchan (hchan any) (? bool)\n&#34;
	&#34;func runtime.newselect (size int) (sel *uint8)\n&#34;
	&#34;func runtime.selectsend (sel *uint8, hchan chan&lt;- any, elem any) (selected bool)\n&#34;
	&#34;func runtime.selectrecv (sel *uint8, hchan &lt;-chan any, elem *any) (selected bool)\n&#34;
	&#34;func runtime.selectdefault (sel *uint8) (selected bool)\n&#34;
	&#34;func runtime.selectgo (sel *uint8)\n&#34;
	&#34;func runtime.makeslice (nel int, cap int, width int) (ary []any)\n&#34;
	&#34;func runtime.sliceslice (old []any, lb int, hb int, width int) (ary []any)\n&#34;
	&#34;func runtime.slicearray (old *any, nel int, lb int, hb int, width int) (ary []any)\n&#34;
	&#34;func runtime.arraytoslice (old *any, nel int) (ary []any)\n&#34;
	&#34;func runtime.closure ()\n&#34;
	&#34;func runtime.int64div (? int64, ? int64) (? int64)\n&#34;
	&#34;func runtime.uint64div (? uint64, ? uint64) (? uint64)\n&#34;
	&#34;func runtime.int64mod (? int64, ? int64) (? int64)\n&#34;
	&#34;func runtime.uint64mod (? uint64, ? uint64) (? uint64)\n&#34;
	&#34;func runtime.float64toint64 (? float64) (? int64)\n&#34;
	&#34;func runtime.int64tofloat64 (? int64) (? float64)\n&#34;
	&#34;\n&#34;
	&#34;$$\n&#34;;
char *unsafeimport =
	&#34;package unsafe\n&#34;
	&#34;type unsafe.Pointer *any\n&#34;
	&#34;func unsafe.Offsetof (? any) (? int)\n&#34;
	&#34;func unsafe.Sizeof (? any) (? int)\n&#34;
	&#34;func unsafe.Alignof (? any) (? int)\n&#34;
	&#34;func unsafe.Typeof (i interface { }) (typ interface { })\n&#34;
	&#34;func unsafe.Reflect (i interface { }) (typ interface { }, addr unsafe.Pointer)\n&#34;
	&#34;func unsafe.Unreflect (typ interface { }, addr unsafe.Pointer) (ret interface { })\n&#34;
	&#34;\n&#34;
	&#34;$$\n&#34;;
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
