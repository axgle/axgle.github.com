<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/ip.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/net/ip.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// IP address manipulations</span>
<a id="L6"></a><span class="comment">//</span>
<a id="L7"></a><span class="comment">// IPv4 addresses are 4 bytes; IPv6 addresses are 16 bytes.</span>
<a id="L8"></a><span class="comment">// An IPv4 address can be converted to an IPv6 address by</span>
<a id="L9"></a><span class="comment">// adding a canonical prefix (10 zeros, 2 0xFFs).</span>
<a id="L10"></a><span class="comment">// This library accepts either size of byte array but always</span>
<a id="L11"></a><span class="comment">// returns 16-byte addresses.</span>

<a id="L13"></a>package net

<a id="L15"></a><span class="comment">// IP address lengths (bytes).</span>
<a id="L16"></a>const (
    <a id="L17"></a>IPv4len = 4;
    <a id="L18"></a>IPv6len = 16;
<a id="L19"></a>)

<a id="L21"></a><span class="comment">// An IP is a single IP address, an array of bytes.</span>
<a id="L22"></a><span class="comment">// Functions in this package accept either 4-byte (IP v4)</span>
<a id="L23"></a><span class="comment">// or 16-byte (IP v6) arrays as input.  Unless otherwise</span>
<a id="L24"></a><span class="comment">// specified, functions in this package always return</span>
<a id="L25"></a><span class="comment">// IP addresses in 16-byte form using the canonical</span>
<a id="L26"></a><span class="comment">// embedding.</span>
<a id="L27"></a><span class="comment">//</span>
<a id="L28"></a><span class="comment">// Note that in this documentation, referring to an</span>
<a id="L29"></a><span class="comment">// IP address as an IPv4 address or an IPv6 address</span>
<a id="L30"></a><span class="comment">// is a semantic property of the address, not just the</span>
<a id="L31"></a><span class="comment">// length of the byte array: a 16-byte array can still</span>
<a id="L32"></a><span class="comment">// be an IPv4 address.</span>
<a id="L33"></a>type IP []byte

<a id="L35"></a><span class="comment">// An IP mask is an IP address.</span>
<a id="L36"></a>type IPMask []byte

<a id="L38"></a><span class="comment">// IPv4 returns the IP address (in 16-byte form) of the</span>
<a id="L39"></a><span class="comment">// IPv4 address a.b.c.d.</span>
<a id="L40"></a>func IPv4(a, b, c, d byte) IP {
    <a id="L41"></a>p := make(IP, IPv6len);
    <a id="L42"></a>for i := 0; i &lt; 10; i++ {
        <a id="L43"></a>p[i] = 0
    <a id="L44"></a>}
    <a id="L45"></a>p[10] = 0xff;
    <a id="L46"></a>p[11] = 0xff;
    <a id="L47"></a>p[12] = a;
    <a id="L48"></a>p[13] = b;
    <a id="L49"></a>p[14] = c;
    <a id="L50"></a>p[15] = d;
    <a id="L51"></a>return p;
<a id="L52"></a>}

<a id="L54"></a><span class="comment">// Well-known IPv4 addresses</span>
<a id="L55"></a>var (
    <a id="L56"></a>IPv4bcast     = IPv4(255, 255, 255, 255); <span class="comment">// broadcast</span>
    <a id="L57"></a>IPv4allsys    = IPv4(224, 0, 0, 1);       <span class="comment">// all systems</span>
    <a id="L58"></a>IPv4allrouter = IPv4(224, 0, 0, 2);       <span class="comment">// all routers</span>
    <a id="L59"></a>IPv4zero      = IPv4(0, 0, 0, 0);         <span class="comment">// all zeros</span>
<a id="L60"></a>)

<a id="L62"></a><span class="comment">// Well-known IPv6 addresses</span>
<a id="L63"></a>var (
    <a id="L64"></a>IPzero = make(IP, IPv6len); <span class="comment">// all zeros</span>
<a id="L65"></a>)

<a id="L67"></a><span class="comment">// Is p all zeros?</span>
<a id="L68"></a>func isZeros(p IP) bool {
    <a id="L69"></a>for i := 0; i &lt; len(p); i++ {
        <a id="L70"></a>if p[i] != 0 {
            <a id="L71"></a>return false
        <a id="L72"></a>}
    <a id="L73"></a>}
    <a id="L74"></a>return true;
<a id="L75"></a>}

<a id="L77"></a><span class="comment">// To4 converts the IPv4 address ip to a 4-byte representation.</span>
<a id="L78"></a><span class="comment">// If ip is not an IPv4 address, To4 returns nil.</span>
<a id="L79"></a>func (ip IP) To4() IP {
    <a id="L80"></a>if len(ip) == IPv4len {
        <a id="L81"></a>return ip
    <a id="L82"></a>}
    <a id="L83"></a>if len(ip) == IPv6len &amp;&amp;
        <a id="L84"></a>isZeros(ip[0:10]) &amp;&amp;
        <a id="L85"></a>ip[10] == 0xff &amp;&amp;
        <a id="L86"></a>ip[11] == 0xff {
        <a id="L87"></a>return ip[12:16]
    <a id="L88"></a>}
    <a id="L89"></a>return nil;
<a id="L90"></a>}

<a id="L92"></a><span class="comment">// To16 converts the IP address ip to a 16-byte representation.</span>
<a id="L93"></a><span class="comment">// If ip is not an IP address (it is the wrong length), To16 returns nil.</span>
<a id="L94"></a>func (ip IP) To16() IP {
    <a id="L95"></a>if len(ip) == IPv4len {
        <a id="L96"></a>return IPv4(ip[0], ip[1], ip[2], ip[3])
    <a id="L97"></a>}
    <a id="L98"></a>if len(ip) == IPv6len {
        <a id="L99"></a>return ip
    <a id="L100"></a>}
    <a id="L101"></a>return nil;
<a id="L102"></a>}

<a id="L104"></a><span class="comment">// Default route masks for IPv4.</span>
<a id="L105"></a>var (
    <a id="L106"></a>classAMask = IPMask(IPv4(0xff, 0, 0, 0));
    <a id="L107"></a>classBMask = IPMask(IPv4(0xff, 0xff, 0, 0));
    <a id="L108"></a>classCMask = IPMask(IPv4(0xff, 0xff, 0xff, 0));
<a id="L109"></a>)

<a id="L111"></a><span class="comment">// DefaultMask returns the default IP mask for the IP address ip.</span>
<a id="L112"></a><span class="comment">// Only IPv4 addresses have default masks; DefaultMask returns</span>
<a id="L113"></a><span class="comment">// nil if ip is not a valid IPv4 address.</span>
<a id="L114"></a>func (ip IP) DefaultMask() IPMask {
    <a id="L115"></a>if ip = ip.To4(); ip == nil {
        <a id="L116"></a>return nil
    <a id="L117"></a>}
    <a id="L118"></a>switch true {
    <a id="L119"></a>case ip[0] &lt; 0x80:
        <a id="L120"></a>return classAMask
    <a id="L121"></a>case ip[0] &lt; 0xC0:
        <a id="L122"></a>return classBMask
    <a id="L123"></a>default:
        <a id="L124"></a>return classCMask
    <a id="L125"></a>}
    <a id="L126"></a>return nil; <span class="comment">// not reached</span>
<a id="L127"></a>}

<a id="L129"></a><span class="comment">// Mask returns the result of masking the IP address ip with mask.</span>
<a id="L130"></a>func (ip IP) Mask(mask IPMask) IP {
    <a id="L131"></a>n := len(ip);
    <a id="L132"></a>if n != len(mask) {
        <a id="L133"></a>return nil
    <a id="L134"></a>}
    <a id="L135"></a>out := make(IP, n);
    <a id="L136"></a>for i := 0; i &lt; n; i++ {
        <a id="L137"></a>out[i] = ip[i] &amp; mask[i]
    <a id="L138"></a>}
    <a id="L139"></a>return out;
<a id="L140"></a>}

<a id="L142"></a><span class="comment">// Convert i to decimal string.</span>
<a id="L143"></a>func itod(i uint) string {
    <a id="L144"></a>if i == 0 {
        <a id="L145"></a>return &#34;0&#34;
    <a id="L146"></a>}

    <a id="L148"></a><span class="comment">// Assemble decimal in reverse order.</span>
    <a id="L149"></a>var b [32]byte;
    <a id="L150"></a>bp := len(b);
    <a id="L151"></a>for ; i &gt; 0; i /= 10 {
        <a id="L152"></a>bp--;
        <a id="L153"></a>b[bp] = byte(i%10) + &#39;0&#39;;
    <a id="L154"></a>}

    <a id="L156"></a>return string(b[bp:len(b)]);
<a id="L157"></a>}

<a id="L159"></a><span class="comment">// Convert i to hexadecimal string.</span>
<a id="L160"></a>func itox(i uint) string {
    <a id="L161"></a>if i == 0 {
        <a id="L162"></a>return &#34;0&#34;
    <a id="L163"></a>}

    <a id="L165"></a><span class="comment">// Assemble hexadecimal in reverse order.</span>
    <a id="L166"></a>var b [32]byte;
    <a id="L167"></a>bp := len(b);
    <a id="L168"></a>for ; i &gt; 0; i /= 16 {
        <a id="L169"></a>bp--;
        <a id="L170"></a>b[bp] = &#34;0123456789abcdef&#34;[byte(i%16)];
    <a id="L171"></a>}

    <a id="L173"></a>return string(b[bp:len(b)]);
<a id="L174"></a>}

<a id="L176"></a><span class="comment">// String returns the string form of the IP address ip.</span>
<a id="L177"></a><span class="comment">// If the address is an IPv4 address, the string representation</span>
<a id="L178"></a><span class="comment">// is dotted decimal (&#34;74.125.19.99&#34;).  Otherwise the representation</span>
<a id="L179"></a><span class="comment">// is IPv6 (&#34;2001:4860:0:2001::68&#34;).</span>
<a id="L180"></a>func (ip IP) String() string {
    <a id="L181"></a>p := ip;

    <a id="L183"></a>if len(ip) == 0 {
        <a id="L184"></a>return &#34;&#34;
    <a id="L185"></a>}

    <a id="L187"></a><span class="comment">// If IPv4, use dotted notation.</span>
    <a id="L188"></a>if p4 := p.To4(); len(p4) == 4 {
        <a id="L189"></a>return itod(uint(p4[0])) + &#34;.&#34; +
            <a id="L190"></a>itod(uint(p4[1])) + &#34;.&#34; +
            <a id="L191"></a>itod(uint(p4[2])) + &#34;.&#34; +
            <a id="L192"></a>itod(uint(p4[3]))
    <a id="L193"></a>}
    <a id="L194"></a>if len(p) != IPv6len {
        <a id="L195"></a>return &#34;?&#34;
    <a id="L196"></a>}

    <a id="L198"></a><span class="comment">// Find longest run of zeros.</span>
    <a id="L199"></a>e0 := -1;
    <a id="L200"></a>e1 := -1;
    <a id="L201"></a>for i := 0; i &lt; 16; i += 2 {
        <a id="L202"></a>j := i;
        <a id="L203"></a>for j &lt; 16 &amp;&amp; p[j] == 0 &amp;&amp; p[j+1] == 0 {
            <a id="L204"></a>j += 2
        <a id="L205"></a>}
        <a id="L206"></a>if j &gt; i &amp;&amp; j-i &gt; e1-e0 {
            <a id="L207"></a>e0 = i;
            <a id="L208"></a>e1 = j;
        <a id="L209"></a>}
    <a id="L210"></a>}

    <a id="L212"></a><span class="comment">// Print with possible :: in place of run of zeros</span>
    <a id="L213"></a>var s string;
    <a id="L214"></a>for i := 0; i &lt; 16; i += 2 {
        <a id="L215"></a>if i == e0 {
            <a id="L216"></a>s += &#34;::&#34;;
            <a id="L217"></a>i = e1;
            <a id="L218"></a>if i &gt;= 16 {
                <a id="L219"></a>break
            <a id="L220"></a>}
        <a id="L221"></a>} else if i &gt; 0 {
            <a id="L222"></a>s += &#34;:&#34;
        <a id="L223"></a>}
        <a id="L224"></a>s += itox((uint(p[i]) &lt;&lt; 8) | uint(p[i+1]));
    <a id="L225"></a>}
    <a id="L226"></a>return s;
<a id="L227"></a>}

<a id="L229"></a><span class="comment">// If mask is a sequence of 1 bits followed by 0 bits,</span>
<a id="L230"></a><span class="comment">// return the number of 1 bits.</span>
<a id="L231"></a>func simpleMaskLength(mask IPMask) int {
    <a id="L232"></a>var i int;
    <a id="L233"></a>for i = 0; i &lt; len(mask); i++ {
        <a id="L234"></a>if mask[i] != 0xFF {
            <a id="L235"></a>break
        <a id="L236"></a>}
    <a id="L237"></a>}
    <a id="L238"></a>n := 8 * i;
    <a id="L239"></a>v := mask[i];
    <a id="L240"></a>for v&amp;0x80 != 0 {
        <a id="L241"></a>n++;
        <a id="L242"></a>v &lt;&lt;= 1;
    <a id="L243"></a>}
    <a id="L244"></a>if v != 0 {
        <a id="L245"></a>return -1
    <a id="L246"></a>}
    <a id="L247"></a>for i++; i &lt; len(mask); i++ {
        <a id="L248"></a>if mask[i] != 0 {
            <a id="L249"></a>return -1
        <a id="L250"></a>}
    <a id="L251"></a>}
    <a id="L252"></a>return n;
<a id="L253"></a>}

<a id="L255"></a><span class="comment">// String returns the string representation of mask.</span>
<a id="L256"></a><span class="comment">// If the mask is in the canonical form--ones followed by zeros--the</span>
<a id="L257"></a><span class="comment">// string representation is just the decimal number of ones.</span>
<a id="L258"></a><span class="comment">// If the mask is in a non-canonical form, it is formatted</span>
<a id="L259"></a><span class="comment">// as an IP address.</span>
<a id="L260"></a>func (mask IPMask) String() string {
    <a id="L261"></a>switch len(mask) {
    <a id="L262"></a>case 4:
        <a id="L263"></a>n := simpleMaskLength(mask);
        <a id="L264"></a>if n &gt;= 0 {
            <a id="L265"></a>return itod(uint(n + (IPv6len-IPv4len)*8))
        <a id="L266"></a>}
    <a id="L267"></a>case 16:
        <a id="L268"></a>n := simpleMaskLength(mask);
        <a id="L269"></a>if n &gt;= 0 {
            <a id="L270"></a>return itod(uint(n))
        <a id="L271"></a>}
    <a id="L272"></a>}
    <a id="L273"></a>return IP(mask).String();
<a id="L274"></a>}

<a id="L276"></a><span class="comment">// Parse IPv4 address (d.d.d.d).</span>
<a id="L277"></a>func parseIPv4(s string) IP {
    <a id="L278"></a>var p [IPv4len]byte;
    <a id="L279"></a>i := 0;
    <a id="L280"></a>for j := 0; j &lt; IPv4len; j++ {
        <a id="L281"></a>if j &gt; 0 {
            <a id="L282"></a>if s[i] != &#39;.&#39; {
                <a id="L283"></a>return nil
            <a id="L284"></a>}
            <a id="L285"></a>i++;
        <a id="L286"></a>}
        <a id="L287"></a>var (
            <a id="L288"></a>n   int;
            <a id="L289"></a>ok  bool;
        <a id="L290"></a>)
        <a id="L291"></a>n, i, ok = dtoi(s, i);
        <a id="L292"></a>if !ok || n &gt; 0xFF {
            <a id="L293"></a>return nil
        <a id="L294"></a>}
        <a id="L295"></a>p[j] = byte(n);
    <a id="L296"></a>}
    <a id="L297"></a>if i != len(s) {
        <a id="L298"></a>return nil
    <a id="L299"></a>}
    <a id="L300"></a>return IPv4(p[0], p[1], p[2], p[3]);
<a id="L301"></a>}

<a id="L303"></a><span class="comment">// Parse IPv6 address.  Many forms.</span>
<a id="L304"></a><span class="comment">// The basic form is a sequence of eight colon-separated</span>
<a id="L305"></a><span class="comment">// 16-bit hex numbers separated by colons,</span>
<a id="L306"></a><span class="comment">// as in 0123:4567:89ab:cdef:0123:4567:89ab:cdef.</span>
<a id="L307"></a><span class="comment">// Two exceptions:</span>
<a id="L308"></a><span class="comment">//	* A run of zeros can be replaced with &#34;::&#34;.</span>
<a id="L309"></a><span class="comment">//	* The last 32 bits can be in IPv4 form.</span>
<a id="L310"></a><span class="comment">// Thus, ::ffff:1.2.3.4 is the IPv4 address 1.2.3.4.</span>
<a id="L311"></a>func parseIPv6(s string) IP {
    <a id="L312"></a>p := make(IP, 16);
    <a id="L313"></a>ellipsis := -1; <span class="comment">// position of ellipsis in p</span>
    <a id="L314"></a>i := 0;         <span class="comment">// index in string s</span>

    <a id="L316"></a><span class="comment">// Might have leading ellipsis</span>
    <a id="L317"></a>if len(s) &gt;= 2 &amp;&amp; s[0] == &#39;:&#39; &amp;&amp; s[1] == &#39;:&#39; {
        <a id="L318"></a>ellipsis = 0;
        <a id="L319"></a>i = 2;
        <a id="L320"></a><span class="comment">// Might be only ellipsis</span>
        <a id="L321"></a>if i == len(s) {
            <a id="L322"></a>return p
        <a id="L323"></a>}
    <a id="L324"></a>}

    <a id="L326"></a><span class="comment">// Loop, parsing hex numbers followed by colon.</span>
    <a id="L327"></a>j := 0;
<a id="L328"></a>L:  for j &lt; IPv6len {
        <a id="L329"></a><span class="comment">// Hex number.</span>
        <a id="L330"></a>n, i1, ok := xtoi(s, i);
        <a id="L331"></a>if !ok || n &gt; 0xFFFF {
            <a id="L332"></a>return nil
        <a id="L333"></a>}

        <a id="L335"></a><span class="comment">// If followed by dot, might be in trailing IPv4.</span>
        <a id="L336"></a>if i1 &lt; len(s) &amp;&amp; s[i1] == &#39;.&#39; {
            <a id="L337"></a>if ellipsis &lt; 0 &amp;&amp; j != IPv6len-IPv4len {
                <a id="L338"></a><span class="comment">// Not the right place.</span>
                <a id="L339"></a>return nil
            <a id="L340"></a>}
            <a id="L341"></a>if j+IPv4len &gt; IPv6len {
                <a id="L342"></a><span class="comment">// Not enough room.</span>
                <a id="L343"></a>return nil
            <a id="L344"></a>}
            <a id="L345"></a>p4 := parseIPv4(s[i:len(s)]);
            <a id="L346"></a>if p4 == nil {
                <a id="L347"></a>return nil
            <a id="L348"></a>}
            <a id="L349"></a>p[j] = p4[12];
            <a id="L350"></a>p[j+1] = p4[13];
            <a id="L351"></a>p[j+2] = p4[14];
            <a id="L352"></a>p[j+3] = p4[15];
            <a id="L353"></a>i = len(s);
            <a id="L354"></a>j += 4;
            <a id="L355"></a>break;
        <a id="L356"></a>}

        <a id="L358"></a><span class="comment">// Save this 16-bit chunk.</span>
        <a id="L359"></a>p[j] = byte(n &gt;&gt; 8);
        <a id="L360"></a>p[j+1] = byte(n);
        <a id="L361"></a>j += 2;

        <a id="L363"></a><span class="comment">// Stop at end of string.</span>
        <a id="L364"></a>i = i1;
        <a id="L365"></a>if i == len(s) {
            <a id="L366"></a>break
        <a id="L367"></a>}

        <a id="L369"></a><span class="comment">// Otherwise must be followed by colon and more.</span>
        <a id="L370"></a>if s[i] != &#39;:&#39; &amp;&amp; i+1 == len(s) {
            <a id="L371"></a>return nil
        <a id="L372"></a>}
        <a id="L373"></a>i++;

        <a id="L375"></a><span class="comment">// Look for ellipsis.</span>
        <a id="L376"></a>if s[i] == &#39;:&#39; {
            <a id="L377"></a>if ellipsis &gt;= 0 { <span class="comment">// already have one</span>
                <a id="L378"></a>return nil
            <a id="L379"></a>}
            <a id="L380"></a>ellipsis = j;
            <a id="L381"></a>if i++; i == len(s) { <span class="comment">// can be at end</span>
                <a id="L382"></a>break
            <a id="L383"></a>}
        <a id="L384"></a>}
    <a id="L385"></a>}

    <a id="L387"></a><span class="comment">// Must have used entire string.</span>
    <a id="L388"></a>if i != len(s) {
        <a id="L389"></a>return nil
    <a id="L390"></a>}

    <a id="L392"></a><span class="comment">// If didn&#39;t parse enough, expand ellipsis.</span>
    <a id="L393"></a>if j &lt; IPv6len {
        <a id="L394"></a>if ellipsis &lt; 0 {
            <a id="L395"></a>return nil
        <a id="L396"></a>}
        <a id="L397"></a>n := IPv6len - j;
        <a id="L398"></a>for k := j - 1; k &gt;= ellipsis; k-- {
            <a id="L399"></a>p[k+n] = p[k]
        <a id="L400"></a>}
        <a id="L401"></a>for k := ellipsis + n - 1; k &gt;= ellipsis; k-- {
            <a id="L402"></a>p[k] = 0
        <a id="L403"></a>}
    <a id="L404"></a>}
    <a id="L405"></a>return p;
<a id="L406"></a>}

<a id="L408"></a><span class="comment">// ParseIP parses s as an IP address, returning the result.</span>
<a id="L409"></a><span class="comment">// The string s can be in dotted decimal (&#34;74.125.19.99&#34;)</span>
<a id="L410"></a><span class="comment">// or IPv6 (&#34;2001:4860:0:2001::68&#34;) form.</span>
<a id="L411"></a><span class="comment">// If s is not a valid textual representation of an IP address,</span>
<a id="L412"></a><span class="comment">// ParseIP returns nil.</span>
<a id="L413"></a>func ParseIP(s string) IP {
    <a id="L414"></a>p := parseIPv4(s);
    <a id="L415"></a>if p != nil {
        <a id="L416"></a>return p
    <a id="L417"></a>}
    <a id="L418"></a>return parseIPv6(s);
<a id="L419"></a>}
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
