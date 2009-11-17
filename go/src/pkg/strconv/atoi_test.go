<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strconv/atoi_test.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/strconv/atoi_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package strconv_test

<a id="L7"></a>import (
    <a id="L8"></a>&#34;os&#34;;
    <a id="L9"></a>&#34;reflect&#34;;
    <a id="L10"></a>. &#34;strconv&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)

<a id="L14"></a>type atoui64Test struct {
    <a id="L15"></a>in  string;
    <a id="L16"></a>out uint64;
    <a id="L17"></a>err os.Error;
<a id="L18"></a>}

<a id="L20"></a>var atoui64tests = []atoui64Test{
    <a id="L21"></a>atoui64Test{&#34;&#34;, 0, os.EINVAL},
    <a id="L22"></a>atoui64Test{&#34;0&#34;, 0, nil},
    <a id="L23"></a>atoui64Test{&#34;1&#34;, 1, nil},
    <a id="L24"></a>atoui64Test{&#34;12345&#34;, 12345, nil},
    <a id="L25"></a>atoui64Test{&#34;012345&#34;, 12345, nil},
    <a id="L26"></a>atoui64Test{&#34;12345x&#34;, 0, os.EINVAL},
    <a id="L27"></a>atoui64Test{&#34;98765432100&#34;, 98765432100, nil},
    <a id="L28"></a>atoui64Test{&#34;18446744073709551615&#34;, 1&lt;&lt;64 - 1, nil},
    <a id="L29"></a>atoui64Test{&#34;18446744073709551616&#34;, 1&lt;&lt;64 - 1, os.ERANGE},
    <a id="L30"></a>atoui64Test{&#34;18446744073709551620&#34;, 1&lt;&lt;64 - 1, os.ERANGE},
<a id="L31"></a>}

<a id="L33"></a>var btoui64tests = []atoui64Test{
    <a id="L34"></a>atoui64Test{&#34;&#34;, 0, os.EINVAL},
    <a id="L35"></a>atoui64Test{&#34;0&#34;, 0, nil},
    <a id="L36"></a>atoui64Test{&#34;1&#34;, 1, nil},
    <a id="L37"></a>atoui64Test{&#34;12345&#34;, 12345, nil},
    <a id="L38"></a>atoui64Test{&#34;012345&#34;, 012345, nil},
    <a id="L39"></a>atoui64Test{&#34;0x12345&#34;, 0x12345, nil},
    <a id="L40"></a>atoui64Test{&#34;0X12345&#34;, 0x12345, nil},
    <a id="L41"></a>atoui64Test{&#34;12345x&#34;, 0, os.EINVAL},
    <a id="L42"></a>atoui64Test{&#34;98765432100&#34;, 98765432100, nil},
    <a id="L43"></a>atoui64Test{&#34;18446744073709551615&#34;, 1&lt;&lt;64 - 1, nil},
    <a id="L44"></a>atoui64Test{&#34;18446744073709551616&#34;, 1&lt;&lt;64 - 1, os.ERANGE},
    <a id="L45"></a>atoui64Test{&#34;18446744073709551620&#34;, 1&lt;&lt;64 - 1, os.ERANGE},
    <a id="L46"></a>atoui64Test{&#34;0xFFFFFFFFFFFFFFFF&#34;, 1&lt;&lt;64 - 1, nil},
    <a id="L47"></a>atoui64Test{&#34;0x10000000000000000&#34;, 1&lt;&lt;64 - 1, os.ERANGE},
    <a id="L48"></a>atoui64Test{&#34;01777777777777777777777&#34;, 1&lt;&lt;64 - 1, nil},
    <a id="L49"></a>atoui64Test{&#34;01777777777777777777778&#34;, 0, os.EINVAL},
    <a id="L50"></a>atoui64Test{&#34;02000000000000000000000&#34;, 1&lt;&lt;64 - 1, os.ERANGE},
    <a id="L51"></a>atoui64Test{&#34;0200000000000000000000&#34;, 1 &lt;&lt; 61, nil},
<a id="L52"></a>}

<a id="L54"></a>type atoi64Test struct {
    <a id="L55"></a>in  string;
    <a id="L56"></a>out int64;
    <a id="L57"></a>err os.Error;
<a id="L58"></a>}

<a id="L60"></a>var atoi64tests = []atoi64Test{
    <a id="L61"></a>atoi64Test{&#34;&#34;, 0, os.EINVAL},
    <a id="L62"></a>atoi64Test{&#34;0&#34;, 0, nil},
    <a id="L63"></a>atoi64Test{&#34;-0&#34;, 0, nil},
    <a id="L64"></a>atoi64Test{&#34;1&#34;, 1, nil},
    <a id="L65"></a>atoi64Test{&#34;-1&#34;, -1, nil},
    <a id="L66"></a>atoi64Test{&#34;12345&#34;, 12345, nil},
    <a id="L67"></a>atoi64Test{&#34;-12345&#34;, -12345, nil},
    <a id="L68"></a>atoi64Test{&#34;012345&#34;, 12345, nil},
    <a id="L69"></a>atoi64Test{&#34;-012345&#34;, -12345, nil},
    <a id="L70"></a>atoi64Test{&#34;98765432100&#34;, 98765432100, nil},
    <a id="L71"></a>atoi64Test{&#34;-98765432100&#34;, -98765432100, nil},
    <a id="L72"></a>atoi64Test{&#34;9223372036854775807&#34;, 1&lt;&lt;63 - 1, nil},
    <a id="L73"></a>atoi64Test{&#34;-9223372036854775807&#34;, -(1&lt;&lt;63 - 1), nil},
    <a id="L74"></a>atoi64Test{&#34;9223372036854775808&#34;, 1&lt;&lt;63 - 1, os.ERANGE},
    <a id="L75"></a>atoi64Test{&#34;-9223372036854775808&#34;, -1 &lt;&lt; 63, nil},
    <a id="L76"></a>atoi64Test{&#34;9223372036854775809&#34;, 1&lt;&lt;63 - 1, os.ERANGE},
    <a id="L77"></a>atoi64Test{&#34;-9223372036854775809&#34;, -1 &lt;&lt; 63, os.ERANGE},
<a id="L78"></a>}

<a id="L80"></a>var btoi64tests = []atoi64Test{
    <a id="L81"></a>atoi64Test{&#34;&#34;, 0, os.EINVAL},
    <a id="L82"></a>atoi64Test{&#34;0&#34;, 0, nil},
    <a id="L83"></a>atoi64Test{&#34;-0&#34;, 0, nil},
    <a id="L84"></a>atoi64Test{&#34;1&#34;, 1, nil},
    <a id="L85"></a>atoi64Test{&#34;-1&#34;, -1, nil},
    <a id="L86"></a>atoi64Test{&#34;12345&#34;, 12345, nil},
    <a id="L87"></a>atoi64Test{&#34;-12345&#34;, -12345, nil},
    <a id="L88"></a>atoi64Test{&#34;012345&#34;, 012345, nil},
    <a id="L89"></a>atoi64Test{&#34;-012345&#34;, -012345, nil},
    <a id="L90"></a>atoi64Test{&#34;0x12345&#34;, 0x12345, nil},
    <a id="L91"></a>atoi64Test{&#34;-0X12345&#34;, -0x12345, nil},
    <a id="L92"></a>atoi64Test{&#34;12345x&#34;, 0, os.EINVAL},
    <a id="L93"></a>atoi64Test{&#34;-12345x&#34;, 0, os.EINVAL},
    <a id="L94"></a>atoi64Test{&#34;98765432100&#34;, 98765432100, nil},
    <a id="L95"></a>atoi64Test{&#34;-98765432100&#34;, -98765432100, nil},
    <a id="L96"></a>atoi64Test{&#34;9223372036854775807&#34;, 1&lt;&lt;63 - 1, nil},
    <a id="L97"></a>atoi64Test{&#34;-9223372036854775807&#34;, -(1&lt;&lt;63 - 1), nil},
    <a id="L98"></a>atoi64Test{&#34;9223372036854775808&#34;, 1&lt;&lt;63 - 1, os.ERANGE},
    <a id="L99"></a>atoi64Test{&#34;-9223372036854775808&#34;, -1 &lt;&lt; 63, nil},
    <a id="L100"></a>atoi64Test{&#34;9223372036854775809&#34;, 1&lt;&lt;63 - 1, os.ERANGE},
    <a id="L101"></a>atoi64Test{&#34;-9223372036854775809&#34;, -1 &lt;&lt; 63, os.ERANGE},
<a id="L102"></a>}

<a id="L104"></a>type atoui32Test struct {
    <a id="L105"></a>in  string;
    <a id="L106"></a>out uint32;
    <a id="L107"></a>err os.Error;
<a id="L108"></a>}

<a id="L110"></a>var atoui32tests = []atoui32Test{
    <a id="L111"></a>atoui32Test{&#34;&#34;, 0, os.EINVAL},
    <a id="L112"></a>atoui32Test{&#34;0&#34;, 0, nil},
    <a id="L113"></a>atoui32Test{&#34;1&#34;, 1, nil},
    <a id="L114"></a>atoui32Test{&#34;12345&#34;, 12345, nil},
    <a id="L115"></a>atoui32Test{&#34;012345&#34;, 12345, nil},
    <a id="L116"></a>atoui32Test{&#34;12345x&#34;, 0, os.EINVAL},
    <a id="L117"></a>atoui32Test{&#34;987654321&#34;, 987654321, nil},
    <a id="L118"></a>atoui32Test{&#34;4294967295&#34;, 1&lt;&lt;32 - 1, nil},
    <a id="L119"></a>atoui32Test{&#34;4294967296&#34;, 1&lt;&lt;32 - 1, os.ERANGE},
<a id="L120"></a>}

<a id="L122"></a>type atoi32Test struct {
    <a id="L123"></a>in  string;
    <a id="L124"></a>out int32;
    <a id="L125"></a>err os.Error;
<a id="L126"></a>}

<a id="L128"></a>var atoi32tests = []atoi32Test{
    <a id="L129"></a>atoi32Test{&#34;&#34;, 0, os.EINVAL},
    <a id="L130"></a>atoi32Test{&#34;0&#34;, 0, nil},
    <a id="L131"></a>atoi32Test{&#34;-0&#34;, 0, nil},
    <a id="L132"></a>atoi32Test{&#34;1&#34;, 1, nil},
    <a id="L133"></a>atoi32Test{&#34;-1&#34;, -1, nil},
    <a id="L134"></a>atoi32Test{&#34;12345&#34;, 12345, nil},
    <a id="L135"></a>atoi32Test{&#34;-12345&#34;, -12345, nil},
    <a id="L136"></a>atoi32Test{&#34;012345&#34;, 12345, nil},
    <a id="L137"></a>atoi32Test{&#34;-012345&#34;, -12345, nil},
    <a id="L138"></a>atoi32Test{&#34;12345x&#34;, 0, os.EINVAL},
    <a id="L139"></a>atoi32Test{&#34;-12345x&#34;, 0, os.EINVAL},
    <a id="L140"></a>atoi32Test{&#34;987654321&#34;, 987654321, nil},
    <a id="L141"></a>atoi32Test{&#34;-987654321&#34;, -987654321, nil},
    <a id="L142"></a>atoi32Test{&#34;2147483647&#34;, 1&lt;&lt;31 - 1, nil},
    <a id="L143"></a>atoi32Test{&#34;-2147483647&#34;, -(1&lt;&lt;31 - 1), nil},
    <a id="L144"></a>atoi32Test{&#34;2147483648&#34;, 1&lt;&lt;31 - 1, os.ERANGE},
    <a id="L145"></a>atoi32Test{&#34;-2147483648&#34;, -1 &lt;&lt; 31, nil},
    <a id="L146"></a>atoi32Test{&#34;2147483649&#34;, 1&lt;&lt;31 - 1, os.ERANGE},
    <a id="L147"></a>atoi32Test{&#34;-2147483649&#34;, -1 &lt;&lt; 31, os.ERANGE},
<a id="L148"></a>}

<a id="L150"></a>func init() {
    <a id="L151"></a><span class="comment">// The atoi routines return NumErrors wrapping</span>
    <a id="L152"></a><span class="comment">// the error and the string.  Convert the tables above.</span>
    <a id="L153"></a>for i := range atoui64tests {
        <a id="L154"></a>test := &amp;atoui64tests[i];
        <a id="L155"></a>if test.err != nil {
            <a id="L156"></a>test.err = &amp;NumError{test.in, test.err}
        <a id="L157"></a>}
    <a id="L158"></a>}
    <a id="L159"></a>for i := range btoui64tests {
        <a id="L160"></a>test := &amp;btoui64tests[i];
        <a id="L161"></a>if test.err != nil {
            <a id="L162"></a>test.err = &amp;NumError{test.in, test.err}
        <a id="L163"></a>}
    <a id="L164"></a>}
    <a id="L165"></a>for i := range atoi64tests {
        <a id="L166"></a>test := &amp;atoi64tests[i];
        <a id="L167"></a>if test.err != nil {
            <a id="L168"></a>test.err = &amp;NumError{test.in, test.err}
        <a id="L169"></a>}
    <a id="L170"></a>}
    <a id="L171"></a>for i := range btoi64tests {
        <a id="L172"></a>test := &amp;btoi64tests[i];
        <a id="L173"></a>if test.err != nil {
            <a id="L174"></a>test.err = &amp;NumError{test.in, test.err}
        <a id="L175"></a>}
    <a id="L176"></a>}
    <a id="L177"></a>for i := range atoui32tests {
        <a id="L178"></a>test := &amp;atoui32tests[i];
        <a id="L179"></a>if test.err != nil {
            <a id="L180"></a>test.err = &amp;NumError{test.in, test.err}
        <a id="L181"></a>}
    <a id="L182"></a>}
    <a id="L183"></a>for i := range atoi32tests {
        <a id="L184"></a>test := &amp;atoi32tests[i];
        <a id="L185"></a>if test.err != nil {
            <a id="L186"></a>test.err = &amp;NumError{test.in, test.err}
        <a id="L187"></a>}
    <a id="L188"></a>}
<a id="L189"></a>}

<a id="L191"></a>func TestAtoui64(t *testing.T) {
    <a id="L192"></a>for i := range atoui64tests {
        <a id="L193"></a>test := &amp;atoui64tests[i];
        <a id="L194"></a>out, err := Atoui64(test.in);
        <a id="L195"></a>if test.out != out || !reflect.DeepEqual(test.err, err) {
            <a id="L196"></a>t.Errorf(&#34;Atoui64(%q) = %v, %v want %v, %v\n&#34;,
                <a id="L197"></a>test.in, out, err, test.out, test.err)
        <a id="L198"></a>}
    <a id="L199"></a>}
<a id="L200"></a>}

<a id="L202"></a>func TestBtoui64(t *testing.T) {
    <a id="L203"></a>for i := range btoui64tests {
        <a id="L204"></a>test := &amp;btoui64tests[i];
        <a id="L205"></a>out, err := Btoui64(test.in, 0);
        <a id="L206"></a>if test.out != out || !reflect.DeepEqual(test.err, err) {
            <a id="L207"></a>t.Errorf(&#34;Btoui64(%q) = %v, %v want %v, %v\n&#34;,
                <a id="L208"></a>test.in, out, err, test.out, test.err)
        <a id="L209"></a>}
    <a id="L210"></a>}
<a id="L211"></a>}

<a id="L213"></a>func TestAtoi64(t *testing.T) {
    <a id="L214"></a>for i := range atoi64tests {
        <a id="L215"></a>test := &amp;atoi64tests[i];
        <a id="L216"></a>out, err := Atoi64(test.in);
        <a id="L217"></a>if test.out != out || !reflect.DeepEqual(test.err, err) {
            <a id="L218"></a>t.Errorf(&#34;Atoi64(%q) = %v, %v want %v, %v\n&#34;,
                <a id="L219"></a>test.in, out, err, test.out, test.err)
        <a id="L220"></a>}
    <a id="L221"></a>}
<a id="L222"></a>}

<a id="L224"></a>func TestBtoi64(t *testing.T) {
    <a id="L225"></a>for i := range btoi64tests {
        <a id="L226"></a>test := &amp;btoi64tests[i];
        <a id="L227"></a>out, err := Btoi64(test.in, 0);
        <a id="L228"></a>if test.out != out || !reflect.DeepEqual(test.err, err) {
            <a id="L229"></a>t.Errorf(&#34;Btoi64(%q) = %v, %v want %v, %v\n&#34;,
                <a id="L230"></a>test.in, out, err, test.out, test.err)
        <a id="L231"></a>}
    <a id="L232"></a>}
<a id="L233"></a>}

<a id="L235"></a>func TestAtoui(t *testing.T) {
    <a id="L236"></a>switch IntSize {
    <a id="L237"></a>case 32:
        <a id="L238"></a>for i := range atoui32tests {
            <a id="L239"></a>test := &amp;atoui32tests[i];
            <a id="L240"></a>out, err := Atoui(test.in);
            <a id="L241"></a>if test.out != uint32(out) || !reflect.DeepEqual(test.err, err) {
                <a id="L242"></a>t.Errorf(&#34;Atoui(%q) = %v, %v want %v, %v\n&#34;,
                    <a id="L243"></a>test.in, out, err, test.out, test.err)
            <a id="L244"></a>}
        <a id="L245"></a>}
    <a id="L246"></a>case 64:
        <a id="L247"></a>for i := range atoui64tests {
            <a id="L248"></a>test := &amp;atoui64tests[i];
            <a id="L249"></a>out, err := Atoui(test.in);
            <a id="L250"></a>if test.out != uint64(out) || !reflect.DeepEqual(test.err, err) {
                <a id="L251"></a>t.Errorf(&#34;Atoui(%q) = %v, %v want %v, %v\n&#34;,
                    <a id="L252"></a>test.in, out, err, test.out, test.err)
            <a id="L253"></a>}
        <a id="L254"></a>}
    <a id="L255"></a>}
<a id="L256"></a>}

<a id="L258"></a>func TestAtoi(t *testing.T) {
    <a id="L259"></a>switch IntSize {
    <a id="L260"></a>case 32:
        <a id="L261"></a>for i := range atoi32tests {
            <a id="L262"></a>test := &amp;atoi32tests[i];
            <a id="L263"></a>out, err := Atoi(test.in);
            <a id="L264"></a>if test.out != int32(out) || !reflect.DeepEqual(test.err, err) {
                <a id="L265"></a>t.Errorf(&#34;Atoi(%q) = %v, %v want %v, %v\n&#34;,
                    <a id="L266"></a>test.in, out, err, test.out, test.err)
            <a id="L267"></a>}
        <a id="L268"></a>}
    <a id="L269"></a>case 64:
        <a id="L270"></a>for i := range atoi64tests {
            <a id="L271"></a>test := &amp;atoi64tests[i];
            <a id="L272"></a>out, err := Atoi(test.in);
            <a id="L273"></a>if test.out != int64(out) || !reflect.DeepEqual(test.err, err) {
                <a id="L274"></a>t.Errorf(&#34;Atoi(%q) = %v, %v want %v, %v\n&#34;,
                    <a id="L275"></a>test.in, out, err, test.out, test.err)
            <a id="L276"></a>}
        <a id="L277"></a>}
    <a id="L278"></a>}
<a id="L279"></a>}
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
