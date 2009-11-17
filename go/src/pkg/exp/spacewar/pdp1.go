<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/spacewar/pdp1.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

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
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/spacewar/pdp1.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright (c) 1996 Barry Silverman, Brian Silverman, Vadim Gerasimov.</span>
<a id="L2"></a><span class="comment">// Portions Copyright (c) 2009 The Go Authors.</span>
<a id="L3"></a><span class="comment">//</span>
<a id="L4"></a><span class="comment">// Permission is hereby granted, free of charge, to any person obtaining a copy</span>
<a id="L5"></a><span class="comment">// of this software and associated documentation files (the &#34;Software&#34;), to deal</span>
<a id="L6"></a><span class="comment">// in the Software without restriction, including without limitation the rights</span>
<a id="L7"></a><span class="comment">// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell</span>
<a id="L8"></a><span class="comment">// copies of the Software, and to permit persons to whom the Software is</span>
<a id="L9"></a><span class="comment">// furnished to do so, subject to the following conditions:</span>
<a id="L10"></a><span class="comment">//</span>
<a id="L11"></a><span class="comment">// The above copyright notice and this permission notice shall be included in</span>
<a id="L12"></a><span class="comment">// all copies or substantial portions of the Software.</span>
<a id="L13"></a><span class="comment">//</span>
<a id="L14"></a><span class="comment">// THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR</span>
<a id="L15"></a><span class="comment">// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,</span>
<a id="L16"></a><span class="comment">// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE</span>
<a id="L17"></a><span class="comment">// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER</span>
<a id="L18"></a><span class="comment">// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,</span>
<a id="L19"></a><span class="comment">// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN</span>
<a id="L20"></a><span class="comment">// THE SOFTWARE.</span>

<a id="L22"></a><span class="comment">// This package and spacewar.go implement a simple PDP-1 emulator</span>
<a id="L23"></a><span class="comment">// complete enough to run the original PDP-1 video game Spacewar!</span>
<a id="L24"></a><span class="comment">//</span>
<a id="L25"></a><span class="comment">// They are a translation of the Java emulator pdp1.java in</span>
<a id="L26"></a><span class="comment">// http://spacewar.oversigma.com/sources/sources.zip.</span>
<a id="L27"></a><span class="comment">//</span>
<a id="L28"></a><span class="comment">// See also the PDP-1 handbook at http://www.dbit.com/~greeng3/pdp1/pdp1.html</span>
<a id="L29"></a><span class="comment">//</span>
<a id="L30"></a><span class="comment">// http://spacewar.oversigma.com/readme.html reads:</span>
<a id="L31"></a><span class="comment">//</span>
<a id="L32"></a><span class="comment">//	Spacewar! was conceived in 1961 by Martin Graetz, Stephen Russell,</span>
<a id="L33"></a><span class="comment">//	and Wayne Wiitanen. It was first realized on the PDP-1 in 1962 by</span>
<a id="L34"></a><span class="comment">//	Stephen Russell, Peter Samson, Dan Edwards, and Martin Graetz,</span>
<a id="L35"></a><span class="comment">//	together with Alan Kotok, Steve Piner, and Robert A Saunders.</span>
<a id="L36"></a><span class="comment">//	Spacewar! is in the public domain, but this credit paragraph must</span>
<a id="L37"></a><span class="comment">//	accompany all distributed versions of the program.</span>
<a id="L38"></a><span class="comment">//</span>
<a id="L39"></a><span class="comment">//	This is the original version! Martin Graetz provided us with a</span>
<a id="L40"></a><span class="comment">//	printed version of the source. We typed in in again - it was about</span>
<a id="L41"></a><span class="comment">//	40 pages long - and re-assembled it with a PDP-1 assembler written</span>
<a id="L42"></a><span class="comment">//	in PERL. The resulting binary runs on a PDP-1 emulator written as</span>
<a id="L43"></a><span class="comment">//	a Java applet. The code is extremely faithful to the original. There</span>
<a id="L44"></a><span class="comment">//	are only two changes. 1)The spaceships have been made bigger and</span>
<a id="L45"></a><span class="comment">//	2) The overall timing has been special cased to deal with varying</span>
<a id="L46"></a><span class="comment">//	machine speeds.</span>
<a id="L47"></a><span class="comment">//</span>
<a id="L48"></a><span class="comment">//	The &#34;a&#34;, &#34;s&#34;, &#34;d&#34;, &#34;f&#34; keys control one of the spaceships. The &#34;k&#34;,</span>
<a id="L49"></a><span class="comment">//	&#34;l&#34;, &#34;;&#34;, &#34;&#39;&#34; keys control the other. The controls are spin one</span>
<a id="L50"></a><span class="comment">//	way, spin the other, thrust, and fire.</span>
<a id="L51"></a><span class="comment">//</span>
<a id="L52"></a><span class="comment">//	Barry Silverman</span>
<a id="L53"></a><span class="comment">//	Brian Silverman</span>
<a id="L54"></a><span class="comment">//	Vadim Gerasimov</span>
<a id="L55"></a><span class="comment">//</span>
<a id="L56"></a>package pdp1

<a id="L58"></a>import (
    <a id="L59"></a>&#34;bufio&#34;;
    <a id="L60"></a>&#34;fmt&#34;;
    <a id="L61"></a>&#34;os&#34;;
    <a id="L62"></a>&#34;io&#34;;
<a id="L63"></a>)

<a id="L65"></a>type Word uint32

<a id="L67"></a>const mask = 0777777
<a id="L68"></a>const sign = 0400000

<a id="L70"></a>const (
    <a id="L71"></a>_   = iota; <span class="comment">// 00</span>
    <a id="L72"></a>opAND;
    <a id="L73"></a>opIOR;
    <a id="L74"></a>opXOR;
    <a id="L75"></a>opXCT;
    <a id="L76"></a>_;
    <a id="L77"></a>_;
    <a id="L78"></a>opCALJDA;

    <a id="L80"></a>opLAC; <span class="comment">// 10</span>
    <a id="L81"></a>opLIO;
    <a id="L82"></a>opDAC;
    <a id="L83"></a>opDAP;
    <a id="L84"></a>_;
    <a id="L85"></a>opDIO;
    <a id="L86"></a>opDZM;
    <a id="L87"></a>_;

    <a id="L89"></a>opADD; <span class="comment">// 20</span>
    <a id="L90"></a>opSUB;
    <a id="L91"></a>opIDX;
    <a id="L92"></a>opISP;
    <a id="L93"></a>opSAD;
    <a id="L94"></a>opSAS;
    <a id="L95"></a>opMUS;
    <a id="L96"></a>opDIS;

    <a id="L98"></a>opJMP; <span class="comment">// 30</span>
    <a id="L99"></a>opJSP;
    <a id="L100"></a>opSKP;
    <a id="L101"></a>opSFT;
    <a id="L102"></a>opLAW;
    <a id="L103"></a>opIOT;
    <a id="L104"></a>_;
    <a id="L105"></a>opOPR;
<a id="L106"></a>)

<a id="L108"></a><span class="comment">// A Trapper represents an object with a Trap method.</span>
<a id="L109"></a><span class="comment">// The machine calls the Trap method to implement the</span>
<a id="L110"></a><span class="comment">// PDP-1 IOT instruction.</span>
<a id="L111"></a>type Trapper interface {
    <a id="L112"></a>Trap(y Word);
<a id="L113"></a>}

<a id="L115"></a><span class="comment">// An M represents the machine state of a PDP-1.</span>
<a id="L116"></a><span class="comment">// Clients can set Display to install an output device.</span>
<a id="L117"></a>type M struct {
    <a id="L118"></a>AC, IO, PC, OV Word;
    <a id="L119"></a>Mem            [010000]Word;
    <a id="L120"></a>Flag           [7]bool;
    <a id="L121"></a>Sense          [7]bool;
    <a id="L122"></a>Halt           bool;
<a id="L123"></a>}


<a id="L126"></a><span class="comment">// Step runs a single machine instruction.</span>
<a id="L127"></a>func (m *M) Step(t Trapper) os.Error {
    <a id="L128"></a>inst := m.Mem[m.PC];
    <a id="L129"></a>m.PC++;
    <a id="L130"></a>return m.run(inst, t);
<a id="L131"></a>}

<a id="L133"></a><span class="comment">// Normalize actual 32-bit integer i to 18-bit ones-complement integer.</span>
<a id="L134"></a><span class="comment">// Interpret mod 0777777, because 0777777 == -0 == +0 == 0000000.</span>
<a id="L135"></a>func norm(i Word) Word {
    <a id="L136"></a>i += i &gt;&gt; 18;
    <a id="L137"></a>i &amp;= mask;
    <a id="L138"></a>if i == mask {
        <a id="L139"></a>i = 0
    <a id="L140"></a>}
    <a id="L141"></a>return i;
<a id="L142"></a>}

<a id="L144"></a>type UnknownInstrError struct {
    <a id="L145"></a>Inst Word;
    <a id="L146"></a>PC   Word;
<a id="L147"></a>}

<a id="L149"></a>func (e UnknownInstrError) String() string {
    <a id="L150"></a>return fmt.Sprintf(&#34;unknown instruction %06o at %06o&#34;, e.Inst, e.PC)
<a id="L151"></a>}

<a id="L153"></a>type HaltError Word

<a id="L155"></a>func (e HaltError) String() string {
    <a id="L156"></a>return fmt.Sprintf(&#34;executed HLT instruction at %06o&#34;, e)
<a id="L157"></a>}

<a id="L159"></a>type LoopError Word

<a id="L161"></a>func (e LoopError) String() string { return fmt.Sprintf(&#34;indirect load looping at %06o&#34;, e) }

<a id="L163"></a>func (m *M) run(inst Word, t Trapper) os.Error {
    <a id="L164"></a>ib, y := (inst&gt;&gt;12)&amp;1, inst&amp;07777;
    <a id="L165"></a>op := inst &gt;&gt; 13;
    <a id="L166"></a>if op &lt; opSKP &amp;&amp; op != opCALJDA {
        <a id="L167"></a>for n := 0; ib != 0; n++ {
            <a id="L168"></a>if n &gt; 07777 {
                <a id="L169"></a>return LoopError(m.PC - 1)
            <a id="L170"></a>}
            <a id="L171"></a>ib = (m.Mem[y] &gt;&gt; 12) &amp; 1;
            <a id="L172"></a>y = m.Mem[y] &amp; 07777;
        <a id="L173"></a>}
    <a id="L174"></a>}

    <a id="L176"></a>switch op {
    <a id="L177"></a>case opAND:
        <a id="L178"></a>m.AC &amp;= m.Mem[y]
    <a id="L179"></a>case opIOR:
        <a id="L180"></a>m.AC |= m.Mem[y]
    <a id="L181"></a>case opXOR:
        <a id="L182"></a>m.AC ^= m.Mem[y]
    <a id="L183"></a>case opXCT:
        <a id="L184"></a>m.run(m.Mem[y], t)
    <a id="L185"></a>case opCALJDA:
        <a id="L186"></a>a := y;
        <a id="L187"></a>if ib == 0 {
            <a id="L188"></a>a = 64
        <a id="L189"></a>}
        <a id="L190"></a>m.Mem[a] = m.AC;
        <a id="L191"></a>m.AC = (m.OV &lt;&lt; 17) + m.PC;
        <a id="L192"></a>m.PC = a + 1;
    <a id="L193"></a>case opLAC:
        <a id="L194"></a>m.AC = m.Mem[y]
    <a id="L195"></a>case opLIO:
        <a id="L196"></a>m.IO = m.Mem[y]
    <a id="L197"></a>case opDAC:
        <a id="L198"></a>m.Mem[y] = m.AC
    <a id="L199"></a>case opDAP:
        <a id="L200"></a>m.Mem[y] = m.Mem[y]&amp;0770000 | m.AC&amp;07777
    <a id="L201"></a>case opDIO:
        <a id="L202"></a>m.Mem[y] = m.IO
    <a id="L203"></a>case opDZM:
        <a id="L204"></a>m.Mem[y] = 0
    <a id="L205"></a>case opADD:
        <a id="L206"></a>m.AC += m.Mem[y];
        <a id="L207"></a>m.OV = m.AC &gt;&gt; 18;
        <a id="L208"></a>m.AC = norm(m.AC);
    <a id="L209"></a>case opSUB:
        <a id="L210"></a>diffSigns := (m.AC^m.Mem[y])&gt;&gt;17 == 1;
        <a id="L211"></a>m.AC += m.Mem[y] ^ mask;
        <a id="L212"></a>m.AC = norm(m.AC);
        <a id="L213"></a>if diffSigns &amp;&amp; m.Mem[y]&gt;&gt;17 == m.AC&gt;&gt;17 {
            <a id="L214"></a>m.OV = 1
        <a id="L215"></a>}
    <a id="L216"></a>case opIDX:
        <a id="L217"></a>m.AC = norm(m.Mem[y] + 1);
        <a id="L218"></a>m.Mem[y] = m.AC;
    <a id="L219"></a>case opISP:
        <a id="L220"></a>m.AC = norm(m.Mem[y] + 1);
        <a id="L221"></a>m.Mem[y] = m.AC;
        <a id="L222"></a>if m.AC&amp;sign == 0 {
            <a id="L223"></a>m.PC++
        <a id="L224"></a>}
    <a id="L225"></a>case opSAD:
        <a id="L226"></a>if m.AC != m.Mem[y] {
            <a id="L227"></a>m.PC++
        <a id="L228"></a>}
    <a id="L229"></a>case opSAS:
        <a id="L230"></a>if m.AC == m.Mem[y] {
            <a id="L231"></a>m.PC++
        <a id="L232"></a>}
    <a id="L233"></a>case opMUS:
        <a id="L234"></a>if m.IO&amp;1 == 1 {
            <a id="L235"></a>m.AC += m.Mem[y];
            <a id="L236"></a>m.AC = norm(m.AC);
        <a id="L237"></a>}
        <a id="L238"></a>m.IO = (m.IO&gt;&gt;1 | m.AC&lt;&lt;17) &amp; mask;
        <a id="L239"></a>m.AC &gt;&gt;= 1;
    <a id="L240"></a>case opDIS:
        <a id="L241"></a>m.AC, m.IO = (m.AC&lt;&lt;1|m.IO&gt;&gt;17)&amp;mask,
            <a id="L242"></a>((m.IO&lt;&lt;1|m.AC&gt;&gt;17)&amp;mask)^1;
        <a id="L243"></a>if m.IO&amp;1 == 1 {
            <a id="L244"></a>m.AC = m.AC + (m.Mem[y] ^ mask)
        <a id="L245"></a>} else {
            <a id="L246"></a>m.AC = m.AC + 1 + m.Mem[y]
        <a id="L247"></a>}
        <a id="L248"></a>m.AC = norm(m.AC);
    <a id="L249"></a>case opJMP:
        <a id="L250"></a>m.PC = y
    <a id="L251"></a>case opJSP:
        <a id="L252"></a>m.AC = (m.OV &lt;&lt; 17) + m.PC;
        <a id="L253"></a>m.PC = y;
    <a id="L254"></a>case opSKP:
        <a id="L255"></a>cond := y&amp;0100 == 0100 &amp;&amp; m.AC == 0 ||
            <a id="L256"></a>y&amp;0200 == 0200 &amp;&amp; m.AC&gt;&gt;17 == 0 ||
            <a id="L257"></a>y&amp;0400 == 0400 &amp;&amp; m.AC&gt;&gt;17 == 1 ||
            <a id="L258"></a>y&amp;01000 == 01000 &amp;&amp; m.OV == 0 ||
            <a id="L259"></a>y&amp;02000 == 02000 &amp;&amp; m.IO&gt;&gt;17 == 0 ||
            <a id="L260"></a>y&amp;7 != 0 &amp;&amp; !m.Flag[y&amp;7] ||
            <a id="L261"></a>y&amp;070 != 0 &amp;&amp; !m.Sense[(y&amp;070)&gt;&gt;3] ||
            <a id="L262"></a>y&amp;070 == 010;
        <a id="L263"></a>if (ib == 0) == cond {
            <a id="L264"></a>m.PC++
        <a id="L265"></a>}
        <a id="L266"></a>if y&amp;01000 == 01000 {
            <a id="L267"></a>m.OV = 0
        <a id="L268"></a>}
    <a id="L269"></a>case opSFT:
        <a id="L270"></a>for count := inst &amp; 0777; count != 0; count &gt;&gt;= 1 {
            <a id="L271"></a>if count&amp;1 == 0 {
                <a id="L272"></a>continue
            <a id="L273"></a>}
            <a id="L274"></a>switch (inst &gt;&gt; 9) &amp; 017 {
            <a id="L275"></a>case 001: <span class="comment">// rotate AC left</span>
                <a id="L276"></a>m.AC = (m.AC&lt;&lt;1 | m.AC&gt;&gt;17) &amp; mask
            <a id="L277"></a>case 002: <span class="comment">// rotate IO left</span>
                <a id="L278"></a>m.IO = (m.IO&lt;&lt;1 | m.IO&gt;&gt;17) &amp; mask
            <a id="L279"></a>case 003: <span class="comment">// rotate AC and IO left.</span>
                <a id="L280"></a>w := uint64(m.AC)&lt;&lt;18 | uint64(m.IO);
                <a id="L281"></a>w = w&lt;&lt;1 | w&gt;&gt;35;
                <a id="L282"></a>m.AC = Word(w&gt;&gt;18) &amp; mask;
                <a id="L283"></a>m.IO = Word(w) &amp; mask;
            <a id="L284"></a>case 005: <span class="comment">// shift AC left (excluding sign bit)</span>
                <a id="L285"></a>m.AC = (m.AC&lt;&lt;1|m.AC&gt;&gt;17)&amp;mask&amp;^sign | m.AC&amp;sign
            <a id="L286"></a>case 006: <span class="comment">// shift IO left (excluding sign bit)</span>
                <a id="L287"></a>m.IO = (m.IO&lt;&lt;1|m.IO&gt;&gt;17)&amp;mask&amp;^sign | m.IO&amp;sign
            <a id="L288"></a>case 007: <span class="comment">// shift AC and IO left (excluding AC&#39;s sign bit)</span>
                <a id="L289"></a>w := uint64(m.AC)&lt;&lt;18 | uint64(m.IO);
                <a id="L290"></a>w = w&lt;&lt;1 | w&gt;&gt;35;
                <a id="L291"></a>m.AC = Word(w&gt;&gt;18)&amp;mask&amp;^sign | m.AC&amp;sign;
                <a id="L292"></a>m.IO = Word(w)&amp;mask&amp;^sign | m.AC&amp;sign;
            <a id="L293"></a>case 011: <span class="comment">// rotate AC right</span>
                <a id="L294"></a>m.AC = (m.AC&gt;&gt;1 | m.AC&lt;&lt;17) &amp; mask
            <a id="L295"></a>case 012: <span class="comment">// rotate IO right</span>
                <a id="L296"></a>m.IO = (m.IO&gt;&gt;1 | m.IO&lt;&lt;17) &amp; mask
            <a id="L297"></a>case 013: <span class="comment">// rotate AC and IO right</span>
                <a id="L298"></a>w := uint64(m.AC)&lt;&lt;18 | uint64(m.IO);
                <a id="L299"></a>w = w&gt;&gt;1 | w&lt;&lt;35;
                <a id="L300"></a>m.AC = Word(w&gt;&gt;18) &amp; mask;
                <a id="L301"></a>m.IO = Word(w) &amp; mask;
            <a id="L302"></a>case 015: <span class="comment">// shift AC right (excluding sign bit)</span>
                <a id="L303"></a>m.AC = m.AC&gt;&gt;1 | m.AC&amp;sign
            <a id="L304"></a>case 016: <span class="comment">// shift IO right (excluding sign bit)</span>
                <a id="L305"></a>m.IO = m.IO&gt;&gt;1 | m.IO&amp;sign
            <a id="L306"></a>case 017: <span class="comment">// shift AC and IO right (excluding AC&#39;s sign bit)</span>
                <a id="L307"></a>w := uint64(m.AC)&lt;&lt;18 | uint64(m.IO);
                <a id="L308"></a>w = w &gt;&gt; 1;
                <a id="L309"></a>m.AC = Word(w&gt;&gt;18) | m.AC&amp;sign;
                <a id="L310"></a>m.IO = Word(w) &amp; mask;
            <a id="L311"></a>default:
                <a id="L312"></a>goto Unknown
            <a id="L313"></a>}
        <a id="L314"></a>}
    <a id="L315"></a>case opLAW:
        <a id="L316"></a>if ib == 0 {
            <a id="L317"></a>m.AC = y
        <a id="L318"></a>} else {
            <a id="L319"></a>m.AC = y ^ mask
        <a id="L320"></a>}
    <a id="L321"></a>case opIOT:
        <a id="L322"></a>t.Trap(y)
    <a id="L323"></a>case opOPR:
        <a id="L324"></a>if y&amp;0200 == 0200 {
            <a id="L325"></a>m.AC = 0
        <a id="L326"></a>}
        <a id="L327"></a>if y&amp;04000 == 04000 {
            <a id="L328"></a>m.IO = 0
        <a id="L329"></a>}
        <a id="L330"></a>if y&amp;01000 == 01000 {
            <a id="L331"></a>m.AC ^= mask
        <a id="L332"></a>}
        <a id="L333"></a>if y&amp;0400 == 0400 {
            <a id="L334"></a>m.PC--;
            <a id="L335"></a>return HaltError(m.PC);
        <a id="L336"></a>}
        <a id="L337"></a>switch i, f := y&amp;7, y&amp;010 == 010; {
        <a id="L338"></a>case i == 7:
            <a id="L339"></a>for i := 2; i &lt; 7; i++ {
                <a id="L340"></a>m.Flag[i] = f
            <a id="L341"></a>}
        <a id="L342"></a>case i &gt;= 2:
            <a id="L343"></a>m.Flag[i] = f
        <a id="L344"></a>}
    <a id="L345"></a>default:
    <a id="L346"></a>Unknown:
        <a id="L347"></a>return UnknownInstrError{inst, m.PC - 1}
    <a id="L348"></a>}
    <a id="L349"></a>return nil;
<a id="L350"></a>}

<a id="L352"></a><span class="comment">// Load loads the machine&#39;s memory from a text input file</span>
<a id="L353"></a><span class="comment">// listing octal address-value pairs, one per line, matching the</span>
<a id="L354"></a><span class="comment">// regular expression ^[ +]([0-7]+)\t([0-7]+).</span>
<a id="L355"></a>func (m *M) Load(r io.Reader) os.Error {
    <a id="L356"></a>b := bufio.NewReader(r);
    <a id="L357"></a>for {
        <a id="L358"></a>line, err := b.ReadString(&#39;\n&#39;);
        <a id="L359"></a>if err != nil {
            <a id="L360"></a>if err != os.EOF {
                <a id="L361"></a>return err
            <a id="L362"></a>}
            <a id="L363"></a>break;
        <a id="L364"></a>}
        <a id="L365"></a><span class="comment">// look for ^[ +]([0-9]+)\t([0-9]+)</span>
        <a id="L366"></a>if line[0] != &#39; &#39; &amp;&amp; line[0] != &#39;+&#39; {
            <a id="L367"></a>continue
        <a id="L368"></a>}
        <a id="L369"></a>i := 1;
        <a id="L370"></a>a := Word(0);
        <a id="L371"></a>for ; i &lt; len(line) &amp;&amp; &#39;0&#39; &lt;= line[i] &amp;&amp; line[i] &lt;= &#39;7&#39;; i++ {
            <a id="L372"></a>a = a*8 + Word(line[i]-&#39;0&#39;)
        <a id="L373"></a>}
        <a id="L374"></a>if i &gt;= len(line) || line[i] != &#39;\t&#39; || i == 1 {
            <a id="L375"></a>continue
        <a id="L376"></a>}
        <a id="L377"></a>v := Word(0);
        <a id="L378"></a>j := i;
        <a id="L379"></a>for i++; i &lt; len(line) &amp;&amp; &#39;0&#39; &lt;= line[i] &amp;&amp; line[i] &lt;= &#39;7&#39;; i++ {
            <a id="L380"></a>v = v*8 + Word(line[i]-&#39;0&#39;)
        <a id="L381"></a>}
        <a id="L382"></a>if i == j {
            <a id="L383"></a>continue
        <a id="L384"></a>}
        <a id="L385"></a>m.Mem[a] = v;
    <a id="L386"></a>}
    <a id="L387"></a>return nil;
<a id="L388"></a>}
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
