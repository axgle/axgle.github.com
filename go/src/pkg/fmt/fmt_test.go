<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/fmt/fmt_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/fmt/fmt_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package fmt_test

<a id="L7"></a>import (
    <a id="L8"></a>. &#34;fmt&#34;;
    <a id="L9"></a>&#34;io&#34;;
    <a id="L10"></a>&#34;math&#34;;
    <a id="L11"></a>&#34;strings&#34;;
    <a id="L12"></a>&#34;testing&#34;;
<a id="L13"></a>)

<a id="L15"></a>func TestFmtInterface(t *testing.T) {
    <a id="L16"></a>var i1 interface{}
    <a id="L17"></a>i1 = &#34;abc&#34;;
    <a id="L18"></a>s := Sprintf(&#34;%s&#34;, i1);
    <a id="L19"></a>if s != &#34;abc&#34; {
        <a id="L20"></a>t.Errorf(`Sprintf(&#34;%%s&#34;, empty(&#34;abc&#34;)) = %q want %q`, s, &#34;abc&#34;)
    <a id="L21"></a>}
<a id="L22"></a>}

<a id="L24"></a>type fmtTest struct {
    <a id="L25"></a>fmt string;
    <a id="L26"></a>val interface{};
    <a id="L27"></a>out string;
<a id="L28"></a>}

<a id="L30"></a>const b32 uint32 = 1&lt;&lt;32 - 1
<a id="L31"></a>const b64 uint64 = 1&lt;&lt;64 - 1

<a id="L33"></a>var array = []int{1, 2, 3, 4, 5}
<a id="L34"></a>var iarray = []interface{}{1, &#34;hello&#34;, 2.5, nil}

<a id="L36"></a>type A struct {
    <a id="L37"></a>i   int;
    <a id="L38"></a>j   uint;
    <a id="L39"></a>s   string;
    <a id="L40"></a>x   []int;
<a id="L41"></a>}

<a id="L43"></a>var b byte

<a id="L45"></a>var fmttests = []fmtTest{
    <a id="L46"></a><span class="comment">// basic string</span>
    <a id="L47"></a>fmtTest{&#34;%s&#34;, &#34;abc&#34;, &#34;abc&#34;},
    <a id="L48"></a>fmtTest{&#34;%x&#34;, &#34;abc&#34;, &#34;616263&#34;},
    <a id="L49"></a>fmtTest{&#34;%x&#34;, &#34;xyz&#34;, &#34;78797a&#34;},
    <a id="L50"></a>fmtTest{&#34;%X&#34;, &#34;xyz&#34;, &#34;78797A&#34;},
    <a id="L51"></a>fmtTest{&#34;%q&#34;, &#34;abc&#34;, `&#34;abc&#34;`},

    <a id="L53"></a><span class="comment">// basic bytes</span>
    <a id="L54"></a>fmtTest{&#34;%s&#34;, strings.Bytes(&#34;abc&#34;), &#34;abc&#34;},
    <a id="L55"></a>fmtTest{&#34;%x&#34;, strings.Bytes(&#34;abc&#34;), &#34;616263&#34;},
    <a id="L56"></a>fmtTest{&#34;% x&#34;, strings.Bytes(&#34;abc&#34;), &#34;61 62 63&#34;},
    <a id="L57"></a>fmtTest{&#34;%x&#34;, strings.Bytes(&#34;xyz&#34;), &#34;78797a&#34;},
    <a id="L58"></a>fmtTest{&#34;%X&#34;, strings.Bytes(&#34;xyz&#34;), &#34;78797A&#34;},
    <a id="L59"></a>fmtTest{&#34;%q&#34;, strings.Bytes(&#34;abc&#34;), `&#34;abc&#34;`},

    <a id="L61"></a><span class="comment">// escaped strings</span>
    <a id="L62"></a>fmtTest{&#34;%#q&#34;, `abc`, &#34;`abc`&#34;},
    <a id="L63"></a>fmtTest{&#34;%#q&#34;, `&#34;`, &#34;`\&#34;`&#34;},
    <a id="L64"></a>fmtTest{&#34;1 %#q&#34;, `\n`, &#34;1 `\\n`&#34;},
    <a id="L65"></a>fmtTest{&#34;2 %#q&#34;, &#34;\n&#34;, `2 &#34;\n&#34;`},
    <a id="L66"></a>fmtTest{&#34;%q&#34;, `&#34;`, `&#34;\&#34;&#34;`},
    <a id="L67"></a>fmtTest{&#34;%q&#34;, &#34;\a\b\f\r\n\t\v&#34;, `&#34;\a\b\f\r\n\t\v&#34;`},
    <a id="L68"></a>fmtTest{&#34;%q&#34;, &#34;abc\xffdef&#34;, `&#34;abc\xffdef&#34;`},
    <a id="L69"></a>fmtTest{&#34;%q&#34;, &#34;\u263a&#34;, `&#34;\u263a&#34;`},
    <a id="L70"></a>fmtTest{&#34;%q&#34;, &#34;\U0010ffff&#34;, `&#34;\U0010ffff&#34;`},

    <a id="L72"></a><span class="comment">// width</span>
    <a id="L73"></a>fmtTest{&#34;%5s&#34;, &#34;abc&#34;, &#34;  abc&#34;},
    <a id="L74"></a>fmtTest{&#34;%-5s&#34;, &#34;abc&#34;, &#34;abc  &#34;},
    <a id="L75"></a>fmtTest{&#34;%05s&#34;, &#34;abc&#34;, &#34;00abc&#34;},

    <a id="L77"></a><span class="comment">// integers</span>
    <a id="L78"></a>fmtTest{&#34;%d&#34;, 12345, &#34;12345&#34;},
    <a id="L79"></a>fmtTest{&#34;%d&#34;, -12345, &#34;-12345&#34;},
    <a id="L80"></a>fmtTest{&#34;%10d&#34;, 12345, &#34;     12345&#34;},
    <a id="L81"></a>fmtTest{&#34;%10d&#34;, -12345, &#34;    -12345&#34;},
    <a id="L82"></a>fmtTest{&#34;%+10d&#34;, 12345, &#34;    +12345&#34;},
    <a id="L83"></a>fmtTest{&#34;%010d&#34;, 12345, &#34;0000012345&#34;},
    <a id="L84"></a>fmtTest{&#34;%010d&#34;, -12345, &#34;-000012345&#34;},
    <a id="L85"></a>fmtTest{&#34;%-10d&#34;, 12345, &#34;12345     &#34;},
    <a id="L86"></a>fmtTest{&#34;%010.3d&#34;, 1, &#34;       001&#34;},
    <a id="L87"></a>fmtTest{&#34;%010.3d&#34;, -1, &#34;      -001&#34;},
    <a id="L88"></a>fmtTest{&#34;%+d&#34;, 12345, &#34;+12345&#34;},
    <a id="L89"></a>fmtTest{&#34;%+d&#34;, -12345, &#34;-12345&#34;},
    <a id="L90"></a>fmtTest{&#34;% d&#34;, 12345, &#34; 12345&#34;},

    <a id="L92"></a><span class="comment">// erroneous formats</span>
    <a id="L93"></a>fmtTest{&#34;&#34;, 2, &#34;?(extra int=2)&#34;},
    <a id="L94"></a>fmtTest{&#34;%d&#34;, &#34;hello&#34;, &#34;%d(string=hello)&#34;},

    <a id="L96"></a><span class="comment">// old test/fmt_test.go</span>
    <a id="L97"></a>fmtTest{&#34;%d&#34;, 1234, &#34;1234&#34;},
    <a id="L98"></a>fmtTest{&#34;%d&#34;, -1234, &#34;-1234&#34;},
    <a id="L99"></a>fmtTest{&#34;%d&#34;, uint(1234), &#34;1234&#34;},
    <a id="L100"></a>fmtTest{&#34;%d&#34;, uint32(b32), &#34;4294967295&#34;},
    <a id="L101"></a>fmtTest{&#34;%d&#34;, uint64(b64), &#34;18446744073709551615&#34;},
    <a id="L102"></a>fmtTest{&#34;%o&#34;, 01234, &#34;1234&#34;},
    <a id="L103"></a>fmtTest{&#34;%#o&#34;, 01234, &#34;01234&#34;},
    <a id="L104"></a>fmtTest{&#34;%o&#34;, uint32(b32), &#34;37777777777&#34;},
    <a id="L105"></a>fmtTest{&#34;%o&#34;, uint64(b64), &#34;1777777777777777777777&#34;},
    <a id="L106"></a>fmtTest{&#34;%x&#34;, 0x1234abcd, &#34;1234abcd&#34;},
    <a id="L107"></a>fmtTest{&#34;%#x&#34;, 0x1234abcd, &#34;0x1234abcd&#34;},
    <a id="L108"></a>fmtTest{&#34;%x&#34;, b32 - 0x1234567, &#34;fedcba98&#34;},
    <a id="L109"></a>fmtTest{&#34;%X&#34;, 0x1234abcd, &#34;1234ABCD&#34;},
    <a id="L110"></a>fmtTest{&#34;%X&#34;, b32 - 0x1234567, &#34;FEDCBA98&#34;},
    <a id="L111"></a>fmtTest{&#34;%#X&#34;, 0, &#34;0X0&#34;},
    <a id="L112"></a>fmtTest{&#34;%x&#34;, b64, &#34;ffffffffffffffff&#34;},
    <a id="L113"></a>fmtTest{&#34;%b&#34;, 7, &#34;111&#34;},
    <a id="L114"></a>fmtTest{&#34;%b&#34;, b64, &#34;1111111111111111111111111111111111111111111111111111111111111111&#34;},
    <a id="L115"></a>fmtTest{&#34;%e&#34;, float64(1), &#34;1.000000e+00&#34;},
    <a id="L116"></a>fmtTest{&#34;%e&#34;, float64(1234.5678e3), &#34;1.234568e+06&#34;},
    <a id="L117"></a>fmtTest{&#34;%e&#34;, float64(1234.5678e-8), &#34;1.234568e-05&#34;},
    <a id="L118"></a>fmtTest{&#34;%e&#34;, float64(-7), &#34;-7.000000e+00&#34;},
    <a id="L119"></a>fmtTest{&#34;%e&#34;, float64(-1e-9), &#34;-1.000000e-09&#34;},
    <a id="L120"></a>fmtTest{&#34;%f&#34;, float64(1234.5678e3), &#34;1234567.800000&#34;},
    <a id="L121"></a>fmtTest{&#34;%f&#34;, float64(1234.5678e-8), &#34;0.000012&#34;},
    <a id="L122"></a>fmtTest{&#34;%f&#34;, float64(-7), &#34;-7.000000&#34;},
    <a id="L123"></a>fmtTest{&#34;%f&#34;, float64(-1e-9), &#34;-0.000000&#34;},
    <a id="L124"></a>fmtTest{&#34;%g&#34;, float64(1234.5678e3), &#34;1.2345678e+06&#34;},
    <a id="L125"></a>fmtTest{&#34;%g&#34;, float32(1234.5678e3), &#34;1.2345678e+06&#34;},
    <a id="L126"></a>fmtTest{&#34;%g&#34;, float64(1234.5678e-8), &#34;1.2345678e-05&#34;},
    <a id="L127"></a>fmtTest{&#34;%g&#34;, float64(-7), &#34;-7&#34;},
    <a id="L128"></a>fmtTest{&#34;%g&#34;, float64(-1e-9), &#34;-1e-09&#34;},
    <a id="L129"></a>fmtTest{&#34;%g&#34;, float32(-1e-9), &#34;-1e-09&#34;},
    <a id="L130"></a>fmtTest{&#34;%E&#34;, float64(1), &#34;1.000000E+00&#34;},
    <a id="L131"></a>fmtTest{&#34;%E&#34;, float64(1234.5678e3), &#34;1.234568E+06&#34;},
    <a id="L132"></a>fmtTest{&#34;%E&#34;, float64(1234.5678e-8), &#34;1.234568E-05&#34;},
    <a id="L133"></a>fmtTest{&#34;%E&#34;, float64(-7), &#34;-7.000000E+00&#34;},
    <a id="L134"></a>fmtTest{&#34;%E&#34;, float64(-1e-9), &#34;-1.000000E-09&#34;},
    <a id="L135"></a>fmtTest{&#34;%G&#34;, float64(1234.5678e3), &#34;1.2345678E+06&#34;},
    <a id="L136"></a>fmtTest{&#34;%G&#34;, float32(1234.5678e3), &#34;1.2345678E+06&#34;},
    <a id="L137"></a>fmtTest{&#34;%G&#34;, float64(1234.5678e-8), &#34;1.2345678E-05&#34;},
    <a id="L138"></a>fmtTest{&#34;%G&#34;, float64(-7), &#34;-7&#34;},
    <a id="L139"></a>fmtTest{&#34;%G&#34;, float64(-1e-9), &#34;-1E-09&#34;},
    <a id="L140"></a>fmtTest{&#34;%G&#34;, float32(-1e-9), &#34;-1E-09&#34;},
    <a id="L141"></a>fmtTest{&#34;%c&#34;, &#39;x&#39;, &#34;x&#34;},
    <a id="L142"></a>fmtTest{&#34;%c&#34;, 0xe4, &#34;ä&#34;},
    <a id="L143"></a>fmtTest{&#34;%c&#34;, 0x672c, &#34;本&#34;},
    <a id="L144"></a>fmtTest{&#34;%c&#34;, &#39;日&#39;, &#34;日&#34;},
    <a id="L145"></a>fmtTest{&#34;%20.8d&#34;, 1234, &#34;            00001234&#34;},
    <a id="L146"></a>fmtTest{&#34;%20.8d&#34;, -1234, &#34;           -00001234&#34;},
    <a id="L147"></a>fmtTest{&#34;%20d&#34;, 1234, &#34;                1234&#34;},
    <a id="L148"></a>fmtTest{&#34;%-20.8d&#34;, 1234, &#34;00001234            &#34;},
    <a id="L149"></a>fmtTest{&#34;%-20.8d&#34;, -1234, &#34;-00001234           &#34;},
    <a id="L150"></a>fmtTest{&#34;%-#20.8x&#34;, 0x1234abc, &#34;0x01234abc          &#34;},
    <a id="L151"></a>fmtTest{&#34;%-#20.8X&#34;, 0x1234abc, &#34;0X01234ABC          &#34;},
    <a id="L152"></a>fmtTest{&#34;%-#20.8o&#34;, 01234, &#34;00001234            &#34;},
    <a id="L153"></a>fmtTest{&#34;%.20b&#34;, 7, &#34;00000000000000000111&#34;},
    <a id="L154"></a>fmtTest{&#34;%20.5s&#34;, &#34;qwertyuiop&#34;, &#34;               qwert&#34;},
    <a id="L155"></a>fmtTest{&#34;%.5s&#34;, &#34;qwertyuiop&#34;, &#34;qwert&#34;},
    <a id="L156"></a>fmtTest{&#34;%-20.5s&#34;, &#34;qwertyuiop&#34;, &#34;qwert               &#34;},
    <a id="L157"></a>fmtTest{&#34;%20c&#34;, &#39;x&#39;, &#34;                   x&#34;},
    <a id="L158"></a>fmtTest{&#34;%-20c&#34;, &#39;x&#39;, &#34;x                   &#34;},
    <a id="L159"></a>fmtTest{&#34;%20.6e&#34;, 1.2345e3, &#34;        1.234500e+03&#34;},
    <a id="L160"></a>fmtTest{&#34;%20.6e&#34;, 1.2345e-3, &#34;        1.234500e-03&#34;},
    <a id="L161"></a>fmtTest{&#34;%20e&#34;, 1.2345e3, &#34;        1.234500e+03&#34;},
    <a id="L162"></a>fmtTest{&#34;%20e&#34;, 1.2345e-3, &#34;        1.234500e-03&#34;},
    <a id="L163"></a>fmtTest{&#34;%20.8e&#34;, 1.2345e3, &#34;      1.23450000e+03&#34;},
    <a id="L164"></a>fmtTest{&#34;%20f&#34;, float64(1.23456789e3), &#34;         1234.567890&#34;},
    <a id="L165"></a>fmtTest{&#34;%20f&#34;, float64(1.23456789e-3), &#34;            0.001235&#34;},
    <a id="L166"></a>fmtTest{&#34;%20f&#34;, float64(12345678901.23456789), &#34;  12345678901.234568&#34;},
    <a id="L167"></a>fmtTest{&#34;%-20f&#34;, float64(1.23456789e3), &#34;1234.567890         &#34;},
    <a id="L168"></a>fmtTest{&#34;%20.8f&#34;, float64(1.23456789e3), &#34;       1234.56789000&#34;},
    <a id="L169"></a>fmtTest{&#34;%20.8f&#34;, float64(1.23456789e-3), &#34;          0.00123457&#34;},
    <a id="L170"></a>fmtTest{&#34;%g&#34;, float64(1.23456789e3), &#34;1234.56789&#34;},
    <a id="L171"></a>fmtTest{&#34;%g&#34;, float64(1.23456789e-3), &#34;0.00123456789&#34;},
    <a id="L172"></a>fmtTest{&#34;%g&#34;, float64(1.23456789e20), &#34;1.23456789e+20&#34;},
    <a id="L173"></a>fmtTest{&#34;%20e&#34;, math.Inf(1), &#34;                +Inf&#34;},
    <a id="L174"></a>fmtTest{&#34;%-20f&#34;, math.Inf(-1), &#34;-Inf                &#34;},
    <a id="L175"></a>fmtTest{&#34;%20g&#34;, math.NaN(), &#34;                 NaN&#34;},

    <a id="L177"></a><span class="comment">// arrays</span>
    <a id="L178"></a>fmtTest{&#34;%v&#34;, array, &#34;[1 2 3 4 5]&#34;},
    <a id="L179"></a>fmtTest{&#34;%v&#34;, iarray, &#34;[1 hello 2.5 &lt;nil&gt;]&#34;},
    <a id="L180"></a>fmtTest{&#34;%v&#34;, &amp;array, &#34;&amp;[1 2 3 4 5]&#34;},
    <a id="L181"></a>fmtTest{&#34;%v&#34;, &amp;iarray, &#34;&amp;[1 hello 2.5 &lt;nil&gt;]&#34;},

    <a id="L183"></a><span class="comment">// structs</span>
    <a id="L184"></a>fmtTest{&#34;%v&#34;, A{1, 2, &#34;a&#34;, []int{1, 2}}, `{1 2 a [1 2]}`},
    <a id="L185"></a>fmtTest{&#34;%+v&#34;, A{1, 2, &#34;a&#34;, []int{1, 2}}, `{i:1 j:2 s:a x:[1 2]}`},

    <a id="L187"></a><span class="comment">// go syntax</span>
    <a id="L188"></a>fmtTest{&#34;%#v&#34;, A{1, 2, &#34;a&#34;, []int{1, 2}}, `fmt_test.A{i:1, j:0x2, s:&#34;a&#34;, x:[]int{1, 2}}`},
    <a id="L189"></a>fmtTest{&#34;%#v&#34;, &amp;b, &#34;(*uint8)(PTR)&#34;},
    <a id="L190"></a>fmtTest{&#34;%#v&#34;, TestFmtInterface, &#34;(func(*testing.T))(PTR)&#34;},
    <a id="L191"></a>fmtTest{&#34;%#v&#34;, make(chan int), &#34;(chan int)(PTR)&#34;},
    <a id="L192"></a>fmtTest{&#34;%#v&#34;, uint64(1&lt;&lt;64 - 1), &#34;0xffffffffffffffff&#34;},
    <a id="L193"></a>fmtTest{&#34;%#v&#34;, 1000000000, &#34;1000000000&#34;},
<a id="L194"></a>}

<a id="L196"></a>func TestSprintf(t *testing.T) {
    <a id="L197"></a>for _, tt := range fmttests {
        <a id="L198"></a>s := Sprintf(tt.fmt, tt.val);
        <a id="L199"></a>if i := strings.Index(s, &#34;0x&#34;); i &gt;= 0 &amp;&amp; strings.Index(tt.out, &#34;PTR&#34;) &gt;= 0 {
            <a id="L200"></a>j := i + 2;
            <a id="L201"></a>for ; j &lt; len(s); j++ {
                <a id="L202"></a>c := s[j];
                <a id="L203"></a>if (c &lt; &#39;0&#39; || c &gt; &#39;9&#39;) &amp;&amp; (c &lt; &#39;a&#39; || c &gt; &#39;f&#39;) {
                    <a id="L204"></a>break
                <a id="L205"></a>}
            <a id="L206"></a>}
            <a id="L207"></a>s = s[0:i] + &#34;PTR&#34; + s[j:len(s)];
        <a id="L208"></a>}
        <a id="L209"></a>if s != tt.out {
            <a id="L210"></a>if _, ok := tt.val.(string); ok {
                <a id="L211"></a><span class="comment">// Don&#39;t requote the already-quoted strings.</span>
                <a id="L212"></a><span class="comment">// It&#39;s too confusing to read the errors.</span>
                <a id="L213"></a>t.Errorf(&#34;Sprintf(%q, %q) = %s want %s&#34;, tt.fmt, tt.val, s, tt.out)
            <a id="L214"></a>} else {
                <a id="L215"></a>t.Errorf(&#34;Sprintf(%q, %v) = %q want %q&#34;, tt.fmt, tt.val, s, tt.out)
            <a id="L216"></a>}
        <a id="L217"></a>}
    <a id="L218"></a>}
<a id="L219"></a>}

<a id="L221"></a>type flagPrinter struct{}

<a id="L223"></a>func (*flagPrinter) Format(f State, c int) {
    <a id="L224"></a>s := &#34;%&#34;;
    <a id="L225"></a>for i := 0; i &lt; 128; i++ {
        <a id="L226"></a>if f.Flag(i) {
            <a id="L227"></a>s += string(i)
        <a id="L228"></a>}
    <a id="L229"></a>}
    <a id="L230"></a>if w, ok := f.Width(); ok {
        <a id="L231"></a>s += Sprintf(&#34;%d&#34;, w)
    <a id="L232"></a>}
    <a id="L233"></a>if p, ok := f.Precision(); ok {
        <a id="L234"></a>s += Sprintf(&#34;.%d&#34;, p)
    <a id="L235"></a>}
    <a id="L236"></a>s += string(c);
    <a id="L237"></a>io.WriteString(f, &#34;[&#34;+s+&#34;]&#34;);
<a id="L238"></a>}

<a id="L240"></a>type flagTest struct {
    <a id="L241"></a>in  string;
    <a id="L242"></a>out string;
<a id="L243"></a>}

<a id="L245"></a>var flagtests = []flagTest{
    <a id="L246"></a>flagTest{&#34;%a&#34;, &#34;[%a]&#34;},
    <a id="L247"></a>flagTest{&#34;%-a&#34;, &#34;[%-a]&#34;},
    <a id="L248"></a>flagTest{&#34;%+a&#34;, &#34;[%+a]&#34;},
    <a id="L249"></a>flagTest{&#34;%#a&#34;, &#34;[%#a]&#34;},
    <a id="L250"></a>flagTest{&#34;% a&#34;, &#34;[% a]&#34;},
    <a id="L251"></a>flagTest{&#34;%0a&#34;, &#34;[%0a]&#34;},
    <a id="L252"></a>flagTest{&#34;%1.2a&#34;, &#34;[%1.2a]&#34;},
    <a id="L253"></a>flagTest{&#34;%-1.2a&#34;, &#34;[%-1.2a]&#34;},
    <a id="L254"></a>flagTest{&#34;%+1.2a&#34;, &#34;[%+1.2a]&#34;},
    <a id="L255"></a>flagTest{&#34;%-+1.2a&#34;, &#34;[%+-1.2a]&#34;},
    <a id="L256"></a>flagTest{&#34;%-+1.2abc&#34;, &#34;[%+-1.2a]bc&#34;},
    <a id="L257"></a>flagTest{&#34;%-1.2abc&#34;, &#34;[%-1.2a]bc&#34;},
<a id="L258"></a>}

<a id="L260"></a>func TestFlagParser(t *testing.T) {
    <a id="L261"></a>var flagprinter flagPrinter;
    <a id="L262"></a>for _, tt := range flagtests {
        <a id="L263"></a>s := Sprintf(tt.in, &amp;flagprinter);
        <a id="L264"></a>if s != tt.out {
            <a id="L265"></a>t.Errorf(&#34;Sprintf(%q, &amp;flagprinter) =&gt; %q, want %q&#34;, tt.in, s, tt.out)
        <a id="L266"></a>}
    <a id="L267"></a>}
<a id="L268"></a>}

<a id="L270"></a>func TestStructPrinter(t *testing.T) {
    <a id="L271"></a>var s struct {
        <a id="L272"></a>a   string;
        <a id="L273"></a>b   string;
        <a id="L274"></a>c   int;
    <a id="L275"></a>}
    <a id="L276"></a>s.a = &#34;abc&#34;;
    <a id="L277"></a>s.b = &#34;def&#34;;
    <a id="L278"></a>s.c = 123;
    <a id="L279"></a>type Test struct {
        <a id="L280"></a>fmt string;
        <a id="L281"></a>out string;
    <a id="L282"></a>}
    <a id="L283"></a>var tests = []Test{
        <a id="L284"></a>Test{&#34;%v&#34;, &#34;{abc def 123}&#34;},
        <a id="L285"></a>Test{&#34;%+v&#34;, &#34;{a:abc b:def c:123}&#34;},
    <a id="L286"></a>};
    <a id="L287"></a>for _, tt := range tests {
        <a id="L288"></a>out := Sprintf(tt.fmt, s);
        <a id="L289"></a>if out != tt.out {
            <a id="L290"></a>t.Errorf(&#34;Sprintf(%q, &amp;s) = %q, want %q&#34;, tt.fmt, out, tt.out)
        <a id="L291"></a>}
    <a id="L292"></a>}
<a id="L293"></a>}

<a id="L295"></a><span class="comment">// Check map printing using substrings so we don&#39;t depend on the print order.</span>
<a id="L296"></a>func presentInMap(s string, a []string, t *testing.T) {
    <a id="L297"></a>for i := 0; i &lt; len(a); i++ {
        <a id="L298"></a>loc := strings.Index(s, a[i]);
        <a id="L299"></a>if loc &lt; 0 {
            <a id="L300"></a>t.Errorf(&#34;map print: expected to find %q in %q&#34;, a[i], s)
        <a id="L301"></a>}
        <a id="L302"></a><span class="comment">// make sure the match ends here</span>
        <a id="L303"></a>loc += len(a[i]);
        <a id="L304"></a>if loc &gt;= len(s) || (s[loc] != &#39; &#39; &amp;&amp; s[loc] != &#39;]&#39;) {
            <a id="L305"></a>t.Errorf(&#34;map print: %q not properly terminated in %q&#34;, a[i], s)
        <a id="L306"></a>}
    <a id="L307"></a>}
<a id="L308"></a>}

<a id="L310"></a>func TestMapPrinter(t *testing.T) {
    <a id="L311"></a>m0 := make(map[int]string);
    <a id="L312"></a>s := Sprint(m0);
    <a id="L313"></a>if s != &#34;map[]&#34; {
        <a id="L314"></a>t.Errorf(&#34;empty map printed as %q not %q&#34;, s, &#34;map[]&#34;)
    <a id="L315"></a>}
    <a id="L316"></a>m1 := map[int]string{1: &#34;one&#34;, 2: &#34;two&#34;, 3: &#34;three&#34;};
    <a id="L317"></a>a := []string{&#34;1:one&#34;, &#34;2:two&#34;, &#34;3:three&#34;};
    <a id="L318"></a>presentInMap(Sprintf(&#34;%v&#34;, m1), a, t);
    <a id="L319"></a>presentInMap(Sprint(m1), a, t);
<a id="L320"></a>}

<a id="L322"></a>func TestEmptyMap(t *testing.T) {
    <a id="L323"></a>const emptyMapStr = &#34;map[]&#34;;
    <a id="L324"></a>var m map[string]int;
    <a id="L325"></a>s := Sprint(m);
    <a id="L326"></a>if s != emptyMapStr {
        <a id="L327"></a>t.Errorf(&#34;nil map printed as %q not %q&#34;, s, emptyMapStr)
    <a id="L328"></a>}
    <a id="L329"></a>m = make(map[string]int);
    <a id="L330"></a>s = Sprint(m);
    <a id="L331"></a>if s != emptyMapStr {
        <a id="L332"></a>t.Errorf(&#34;empty map printed as %q not %q&#34;, s, emptyMapStr)
    <a id="L333"></a>}
<a id="L334"></a>}
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
