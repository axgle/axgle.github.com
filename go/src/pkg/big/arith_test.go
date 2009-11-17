<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/big/arith_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/big/arith_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package big

<a id="L7"></a>import &#34;testing&#34;


<a id="L10"></a>type funWW func(x, y, c Word) (z1, z0 Word)
<a id="L11"></a>type argWW struct {
    <a id="L12"></a>x, y, c, z1, z0 Word;
<a id="L13"></a>}

<a id="L15"></a>var sumWW = []argWW{
    <a id="L16"></a>argWW{0, 0, 0, 0, 0},
    <a id="L17"></a>argWW{0, 1, 0, 0, 1},
    <a id="L18"></a>argWW{0, 0, 1, 0, 1},
    <a id="L19"></a>argWW{0, 1, 1, 0, 2},
    <a id="L20"></a>argWW{12345, 67890, 0, 0, 80235},
    <a id="L21"></a>argWW{12345, 67890, 1, 0, 80236},
    <a id="L22"></a>argWW{_M, 1, 0, 1, 0},
    <a id="L23"></a>argWW{_M, 0, 1, 1, 0},
    <a id="L24"></a>argWW{_M, 1, 1, 1, 1},
    <a id="L25"></a>argWW{_M, _M, 0, 1, _M - 1},
    <a id="L26"></a>argWW{_M, _M, 1, 1, _M},
<a id="L27"></a>}


<a id="L30"></a>func testFunWW(t *testing.T, msg string, f funWW, a argWW) {
    <a id="L31"></a>z1, z0 := f(a.x, a.y, a.c);
    <a id="L32"></a>if z1 != a.z1 || z0 != a.z0 {
        <a id="L33"></a>t.Errorf(&#34;%s%+v\n\tgot z1:z0 = %#x:%#x; want %#x:%#x&#34;, msg, a, z1, z0, a.z1, a.z0)
    <a id="L34"></a>}
<a id="L35"></a>}


<a id="L38"></a>func TestFunWW(t *testing.T) {
    <a id="L39"></a>for _, a := range sumWW {
        <a id="L40"></a>arg := a;
        <a id="L41"></a>testFunWW(t, &#34;addWW_g&#34;, addWW_g, arg);

        <a id="L43"></a>arg = argWW{a.y, a.x, a.c, a.z1, a.z0};
        <a id="L44"></a>testFunWW(t, &#34;addWW_g symmetric&#34;, addWW_g, arg);

        <a id="L46"></a>arg = argWW{a.z0, a.x, a.c, a.z1, a.y};
        <a id="L47"></a>testFunWW(t, &#34;subWW_g&#34;, subWW_g, arg);

        <a id="L49"></a>arg = argWW{a.z0, a.y, a.c, a.z1, a.x};
        <a id="L50"></a>testFunWW(t, &#34;subWW_g symmetric&#34;, subWW_g, arg);
    <a id="L51"></a>}
<a id="L52"></a>}


<a id="L55"></a>func addr(x []Word) *Word {
    <a id="L56"></a>if len(x) == 0 {
        <a id="L57"></a>return nil
    <a id="L58"></a>}
    <a id="L59"></a>return &amp;x[0];
<a id="L60"></a>}


<a id="L63"></a>type funVV func(z, x, y *Word, n int) (c Word)
<a id="L64"></a>type argVV struct {
    <a id="L65"></a>z, x, y []Word;
    <a id="L66"></a>c       Word;
<a id="L67"></a>}

<a id="L69"></a>var sumVV = []argVV{
    <a id="L70"></a>argVV{},
    <a id="L71"></a>argVV{[]Word{0}, []Word{0}, []Word{0}, 0},
    <a id="L72"></a>argVV{[]Word{1}, []Word{1}, []Word{0}, 0},
    <a id="L73"></a>argVV{[]Word{0}, []Word{_M}, []Word{1}, 1},
    <a id="L74"></a>argVV{[]Word{80235}, []Word{12345}, []Word{67890}, 0},
    <a id="L75"></a>argVV{[]Word{_M - 1}, []Word{_M}, []Word{_M}, 1},
    <a id="L76"></a>argVV{[]Word{0, 0, 0, 0}, []Word{_M, _M, _M, _M}, []Word{1, 0, 0, 0}, 1},
    <a id="L77"></a>argVV{[]Word{0, 0, 0, _M}, []Word{_M, _M, _M, _M - 1}, []Word{1, 0, 0, 0}, 0},
    <a id="L78"></a>argVV{[]Word{0, 0, 0, 0}, []Word{_M, 0, _M, 0}, []Word{1, _M, 0, _M}, 1},
<a id="L79"></a>}


<a id="L82"></a>func testFunVV(t *testing.T, msg string, f funVV, a argVV) {
    <a id="L83"></a>n := len(a.z);
    <a id="L84"></a>z := make([]Word, n);
    <a id="L85"></a>c := f(addr(z), addr(a.x), addr(a.y), n);
    <a id="L86"></a>for i, zi := range z {
        <a id="L87"></a>if zi != a.z[i] {
            <a id="L88"></a>t.Errorf(&#34;%s%+v\n\tgot z[%d] = %#x; want %#x&#34;, msg, a, i, zi, a.z[i]);
            <a id="L89"></a>break;
        <a id="L90"></a>}
    <a id="L91"></a>}
    <a id="L92"></a>if c != a.c {
        <a id="L93"></a>t.Errorf(&#34;%s%+v\n\tgot c = %#x; want %#x&#34;, msg, a, c, a.c)
    <a id="L94"></a>}
<a id="L95"></a>}


<a id="L98"></a>func TestFunVV(t *testing.T) {
    <a id="L99"></a>for _, a := range sumVV {
        <a id="L100"></a>arg := a;
        <a id="L101"></a>testFunVV(t, &#34;addVV_g&#34;, addVV_g, arg);
        <a id="L102"></a>testFunVV(t, &#34;addVV&#34;, addVV, arg);

        <a id="L104"></a>arg = argVV{a.z, a.y, a.x, a.c};
        <a id="L105"></a>testFunVV(t, &#34;addVV_g symmetric&#34;, addVV_g, arg);
        <a id="L106"></a>testFunVV(t, &#34;addVV symmetric&#34;, addVV, arg);

        <a id="L108"></a>arg = argVV{a.x, a.z, a.y, a.c};
        <a id="L109"></a>testFunVV(t, &#34;subVV_g&#34;, subVV_g, arg);
        <a id="L110"></a>testFunVV(t, &#34;subVV&#34;, subVV, arg);

        <a id="L112"></a>arg = argVV{a.y, a.z, a.x, a.c};
        <a id="L113"></a>testFunVV(t, &#34;subVV_g symmetric&#34;, subVV_g, arg);
        <a id="L114"></a>testFunVV(t, &#34;subVV symmetric&#34;, subVV, arg);
    <a id="L115"></a>}
<a id="L116"></a>}


<a id="L119"></a>type funVW func(z, x *Word, y Word, n int) (c Word)
<a id="L120"></a>type argVW struct {
    <a id="L121"></a>z, x []Word;
    <a id="L122"></a>y    Word;
    <a id="L123"></a>c    Word;
<a id="L124"></a>}

<a id="L126"></a>var sumVW = []argVW{
    <a id="L127"></a>argVW{},
    <a id="L128"></a>argVW{[]Word{0}, []Word{0}, 0, 0},
    <a id="L129"></a>argVW{[]Word{1}, []Word{0}, 1, 0},
    <a id="L130"></a>argVW{[]Word{1}, []Word{1}, 0, 0},
    <a id="L131"></a>argVW{[]Word{0}, []Word{_M}, 1, 1},
    <a id="L132"></a>argVW{[]Word{0, 0, 0, 0}, []Word{_M, _M, _M, _M}, 1, 1},
<a id="L133"></a>}

<a id="L135"></a>var prodVW = []argVW{
    <a id="L136"></a>argVW{},
    <a id="L137"></a>argVW{[]Word{0}, []Word{0}, 0, 0},
    <a id="L138"></a>argVW{[]Word{0}, []Word{_M}, 0, 0},
    <a id="L139"></a>argVW{[]Word{0}, []Word{0}, _M, 0},
    <a id="L140"></a>argVW{[]Word{1}, []Word{1}, 1, 0},
    <a id="L141"></a>argVW{[]Word{22793}, []Word{991}, 23, 0},
    <a id="L142"></a>argVW{[]Word{0, 0, 0, 22793}, []Word{0, 0, 0, 991}, 23, 0},
    <a id="L143"></a>argVW{[]Word{0, 0, 0, 0}, []Word{7893475, 7395495, 798547395, 68943}, 0, 0},
    <a id="L144"></a>argVW{[]Word{0, 0, 0, 0}, []Word{0, 0, 0, 0}, 894375984, 0},
    <a id="L145"></a>argVW{[]Word{_M &lt;&lt; 1 &amp; _M}, []Word{_M}, 1 &lt;&lt; 1, _M &gt;&gt; (_W - 1)},
    <a id="L146"></a>argVW{[]Word{_M &lt;&lt; 7 &amp; _M}, []Word{_M}, 1 &lt;&lt; 7, _M &gt;&gt; (_W - 7)},
    <a id="L147"></a>argVW{[]Word{_M &lt;&lt; 7 &amp; _M, _M, _M, _M}, []Word{_M, _M, _M, _M}, 1 &lt;&lt; 7, _M &gt;&gt; (_W - 7)},
<a id="L148"></a>}


<a id="L151"></a>func testFunVW(t *testing.T, msg string, f funVW, a argVW) {
    <a id="L152"></a>n := len(a.z);
    <a id="L153"></a>z := make([]Word, n);
    <a id="L154"></a>c := f(addr(z), addr(a.x), a.y, n);
    <a id="L155"></a>for i, zi := range z {
        <a id="L156"></a>if zi != a.z[i] {
            <a id="L157"></a>t.Errorf(&#34;%s%+v\n\tgot z[%d] = %#x; want %#x&#34;, msg, a, i, zi, a.z[i]);
            <a id="L158"></a>break;
        <a id="L159"></a>}
    <a id="L160"></a>}
    <a id="L161"></a>if c != a.c {
        <a id="L162"></a>t.Errorf(&#34;%s%+v\n\tgot c = %#x; want %#x&#34;, msg, a, c, a.c)
    <a id="L163"></a>}
<a id="L164"></a>}


<a id="L167"></a>func TestFunVW(t *testing.T) {
    <a id="L168"></a>for _, a := range sumVW {
        <a id="L169"></a>arg := a;
        <a id="L170"></a>testFunVW(t, &#34;addVW_g&#34;, addVW_g, arg);
        <a id="L171"></a>testFunVW(t, &#34;addVW&#34;, addVW, arg);

        <a id="L173"></a>arg = argVW{a.x, a.z, a.y, a.c};
        <a id="L174"></a>testFunVW(t, &#34;subVW_g&#34;, subVW_g, arg);
        <a id="L175"></a>testFunVW(t, &#34;subVW&#34;, subVW, arg);
    <a id="L176"></a>}
<a id="L177"></a>}


<a id="L180"></a>type funVWW func(z, x *Word, y, r Word, n int) (c Word)
<a id="L181"></a>type argVWW struct {
    <a id="L182"></a>z, x []Word;
    <a id="L183"></a>y, r Word;
    <a id="L184"></a>c    Word;
<a id="L185"></a>}

<a id="L187"></a>var prodVWW = []argVWW{
    <a id="L188"></a>argVWW{},
    <a id="L189"></a>argVWW{[]Word{0}, []Word{0}, 0, 0, 0},
    <a id="L190"></a>argVWW{[]Word{991}, []Word{0}, 0, 991, 0},
    <a id="L191"></a>argVWW{[]Word{0}, []Word{_M}, 0, 0, 0},
    <a id="L192"></a>argVWW{[]Word{991}, []Word{_M}, 0, 991, 0},
    <a id="L193"></a>argVWW{[]Word{0}, []Word{0}, _M, 0, 0},
    <a id="L194"></a>argVWW{[]Word{991}, []Word{0}, _M, 991, 0},
    <a id="L195"></a>argVWW{[]Word{1}, []Word{1}, 1, 0, 0},
    <a id="L196"></a>argVWW{[]Word{992}, []Word{1}, 1, 991, 0},
    <a id="L197"></a>argVWW{[]Word{22793}, []Word{991}, 23, 0, 0},
    <a id="L198"></a>argVWW{[]Word{22800}, []Word{991}, 23, 7, 0},
    <a id="L199"></a>argVWW{[]Word{0, 0, 0, 22793}, []Word{0, 0, 0, 991}, 23, 0, 0},
    <a id="L200"></a>argVWW{[]Word{7, 0, 0, 22793}, []Word{0, 0, 0, 991}, 23, 7, 0},
    <a id="L201"></a>argVWW{[]Word{0, 0, 0, 0}, []Word{7893475, 7395495, 798547395, 68943}, 0, 0, 0},
    <a id="L202"></a>argVWW{[]Word{991, 0, 0, 0}, []Word{7893475, 7395495, 798547395, 68943}, 0, 991, 0},
    <a id="L203"></a>argVWW{[]Word{0, 0, 0, 0}, []Word{0, 0, 0, 0}, 894375984, 0, 0},
    <a id="L204"></a>argVWW{[]Word{991, 0, 0, 0}, []Word{0, 0, 0, 0}, 894375984, 991, 0},
    <a id="L205"></a>argVWW{[]Word{_M &lt;&lt; 1 &amp; _M}, []Word{_M}, 1 &lt;&lt; 1, 0, _M &gt;&gt; (_W - 1)},
    <a id="L206"></a>argVWW{[]Word{_M&lt;&lt;1&amp;_M + 1}, []Word{_M}, 1 &lt;&lt; 1, 1, _M &gt;&gt; (_W - 1)},
    <a id="L207"></a>argVWW{[]Word{_M &lt;&lt; 7 &amp; _M}, []Word{_M}, 1 &lt;&lt; 7, 0, _M &gt;&gt; (_W - 7)},
    <a id="L208"></a>argVWW{[]Word{_M&lt;&lt;7&amp;_M + 1&lt;&lt;6}, []Word{_M}, 1 &lt;&lt; 7, 1 &lt;&lt; 6, _M &gt;&gt; (_W - 7)},
    <a id="L209"></a>argVWW{[]Word{_M &lt;&lt; 7 &amp; _M, _M, _M, _M}, []Word{_M, _M, _M, _M}, 1 &lt;&lt; 7, 0, _M &gt;&gt; (_W - 7)},
    <a id="L210"></a>argVWW{[]Word{_M&lt;&lt;7&amp;_M + 1&lt;&lt;6, _M, _M, _M}, []Word{_M, _M, _M, _M}, 1 &lt;&lt; 7, 1 &lt;&lt; 6, _M &gt;&gt; (_W - 7)},
<a id="L211"></a>}


<a id="L214"></a>func testFunVWW(t *testing.T, msg string, f funVWW, a argVWW) {
    <a id="L215"></a>n := len(a.z);
    <a id="L216"></a>z := make([]Word, n);
    <a id="L217"></a>c := f(addr(z), addr(a.x), a.y, a.r, n);
    <a id="L218"></a>for i, zi := range z {
        <a id="L219"></a>if zi != a.z[i] {
            <a id="L220"></a>t.Errorf(&#34;%s%+v\n\tgot z[%d] = %#x; want %#x&#34;, msg, a, i, zi, a.z[i]);
            <a id="L221"></a>break;
        <a id="L222"></a>}
    <a id="L223"></a>}
    <a id="L224"></a>if c != a.c {
        <a id="L225"></a>t.Errorf(&#34;%s%+v\n\tgot c = %#x; want %#x&#34;, msg, a, c, a.c)
    <a id="L226"></a>}
<a id="L227"></a>}


<a id="L230"></a><span class="comment">// TODO(gri) mulAddVWW and divWVW are symmetric operations but</span>
<a id="L231"></a><span class="comment">//           their signature is not symmetric. Try to unify.</span>

<a id="L233"></a>type funWVW func(z *Word, xn Word, x *Word, y Word, n int) (r Word)
<a id="L234"></a>type argWVW struct {
    <a id="L235"></a>z   []Word;
    <a id="L236"></a>xn  Word;
    <a id="L237"></a>x   []Word;
    <a id="L238"></a>y   Word;
    <a id="L239"></a>r   Word;
<a id="L240"></a>}

<a id="L242"></a>func testFunWVW(t *testing.T, msg string, f funWVW, a argWVW) {
    <a id="L243"></a>n := len(a.z);
    <a id="L244"></a>z := make([]Word, n);
    <a id="L245"></a>r := f(addr(z), a.xn, addr(a.x), a.y, n);
    <a id="L246"></a>for i, zi := range z {
        <a id="L247"></a>if zi != a.z[i] {
            <a id="L248"></a>t.Errorf(&#34;%s%+v\n\tgot z[%d] = %#x; want %#x&#34;, msg, a, i, zi, a.z[i]);
            <a id="L249"></a>break;
        <a id="L250"></a>}
    <a id="L251"></a>}
    <a id="L252"></a>if r != a.r {
        <a id="L253"></a>t.Errorf(&#34;%s%+v\n\tgot r = %#x; want %#x&#34;, msg, a, r, a.r)
    <a id="L254"></a>}
<a id="L255"></a>}


<a id="L258"></a>func TestFunVWW(t *testing.T) {
    <a id="L259"></a>for _, a := range prodVWW {
        <a id="L260"></a>arg := a;
        <a id="L261"></a>testFunVWW(t, &#34;mulAddVWW_g&#34;, mulAddVWW_g, arg);
        <a id="L262"></a>testFunVWW(t, &#34;mulAddVWW&#34;, mulAddVWW, arg);

        <a id="L264"></a>if a.y != 0 &amp;&amp; a.r &lt; a.y {
            <a id="L265"></a>arg := argWVW{a.x, a.c, a.z, a.y, a.r};
            <a id="L266"></a>testFunWVW(t, &#34;divWVW_g&#34;, divWVW_g, arg);
            <a id="L267"></a>testFunWVW(t, &#34;divWVW&#34;, divWVW, arg);
        <a id="L268"></a>}
    <a id="L269"></a>}
<a id="L270"></a>}


<a id="L273"></a>type mulWWTest struct {
    <a id="L274"></a>x, y Word;
    <a id="L275"></a>q, r Word;
<a id="L276"></a>}


<a id="L279"></a>var mulWWTests = []mulWWTest{
    <a id="L280"></a>mulWWTest{_M, _M, _M - 1, 1},
    <a id="L281"></a><span class="comment">// 32 bit only: mulWWTest{0xc47dfa8c, 50911, 0x98a4, 0x998587f4},</span>
<a id="L282"></a>}


<a id="L285"></a>func TestMulWW(t *testing.T) {
    <a id="L286"></a>for i, test := range mulWWTests {
        <a id="L287"></a>q, r := mulWW_g(test.x, test.y);
        <a id="L288"></a>if q != test.q || r != test.r {
            <a id="L289"></a>t.Errorf(&#34;#%d got (%x, %x) want (%x, %x)&#34;, i, q, r, test.q, test.r)
        <a id="L290"></a>}
    <a id="L291"></a>}
<a id="L292"></a>}


<a id="L295"></a>type mulAddWWWTest struct {
    <a id="L296"></a>x, y, c Word;
    <a id="L297"></a>q, r    Word;
<a id="L298"></a>}


<a id="L301"></a>var mulAddWWWTests = []mulAddWWWTest{
    <a id="L302"></a><span class="comment">// TODO(agl): These will only work on 64-bit platforms.</span>
    <a id="L303"></a><span class="comment">// mulAddWWWTest{15064310297182388543, 0xe7df04d2d35d5d80, 13537600649892366549, 13644450054494335067, 10832252001440893781},</span>
    <a id="L304"></a><span class="comment">// mulAddWWWTest{15064310297182388543, 0xdab2f18048baa68d, 13644450054494335067, 12869334219691522700, 14233854684711418382},</span>
    <a id="L305"></a>mulAddWWWTest{_M, _M, 0, _M - 1, 1},
    <a id="L306"></a>mulAddWWWTest{_M, _M, _M, _M, 0},
<a id="L307"></a>}


<a id="L310"></a>func TestMulAddWWW(t *testing.T) {
    <a id="L311"></a>for i, test := range mulAddWWWTests {
        <a id="L312"></a>q, r := mulAddWWW_g(test.x, test.y, test.c);
        <a id="L313"></a>if q != test.q || r != test.r {
            <a id="L314"></a>t.Errorf(&#34;#%d got (%x, %x) want (%x, %x)&#34;, i, q, r, test.q, test.r)
        <a id="L315"></a>}
    <a id="L316"></a>}
<a id="L317"></a>}
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
