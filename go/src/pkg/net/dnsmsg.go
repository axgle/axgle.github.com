<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/dnsmsg.go</title>

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
	<li>Thu Nov 12 16:00:00 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/net/dnsmsg.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// DNS packet assembly.  See RFC 1035.</span>
<a id="L6"></a><span class="comment">//</span>
<a id="L7"></a><span class="comment">// This is intended to support name resolution during net.Dial.</span>
<a id="L8"></a><span class="comment">// It doesn&#39;t have to be blazing fast.</span>
<a id="L9"></a><span class="comment">//</span>
<a id="L10"></a><span class="comment">// Rather than write the usual handful of routines to pack and</span>
<a id="L11"></a><span class="comment">// unpack every message that can appear on the wire, we use</span>
<a id="L12"></a><span class="comment">// reflection to write a generic pack/unpack for structs and then</span>
<a id="L13"></a><span class="comment">// use it.  Thus, if in the future we need to define new message</span>
<a id="L14"></a><span class="comment">// structs, no new pack/unpack/printing code needs to be written.</span>
<a id="L15"></a><span class="comment">//</span>
<a id="L16"></a><span class="comment">// The first half of this file defines the DNS message formats.</span>
<a id="L17"></a><span class="comment">// The second half implements the conversion to and from wire format.</span>
<a id="L18"></a><span class="comment">// A few of the structure elements have string tags to aid the</span>
<a id="L19"></a><span class="comment">// generic pack/unpack routines.</span>
<a id="L20"></a><span class="comment">//</span>
<a id="L21"></a><span class="comment">// TODO(rsc)  There are enough names defined in this file that they&#39;re all</span>
<a id="L22"></a><span class="comment">// prefixed with _DNS_.  Perhaps put this in its own package later.</span>

<a id="L24"></a>package net

<a id="L26"></a>import (
    <a id="L27"></a>&#34;fmt&#34;;
    <a id="L28"></a>&#34;os&#34;;
    <a id="L29"></a>&#34;reflect&#34;;
<a id="L30"></a>)

<a id="L32"></a><span class="comment">// Packet formats</span>

<a id="L34"></a><span class="comment">// Wire constants.</span>
<a id="L35"></a>const (
    <a id="L36"></a><span class="comment">// valid _DNS_RR_Header.Rrtype and _DNS_Question.qtype</span>
    <a id="L37"></a>_DNS_TypeA     = 1;
    <a id="L38"></a>_DNS_TypeNS    = 2;
    <a id="L39"></a>_DNS_TypeMD    = 3;
    <a id="L40"></a>_DNS_TypeMF    = 4;
    <a id="L41"></a>_DNS_TypeCNAME = 5;
    <a id="L42"></a>_DNS_TypeSOA   = 6;
    <a id="L43"></a>_DNS_TypeMB    = 7;
    <a id="L44"></a>_DNS_TypeMG    = 8;
    <a id="L45"></a>_DNS_TypeMR    = 9;
    <a id="L46"></a>_DNS_TypeNULL  = 10;
    <a id="L47"></a>_DNS_TypeWKS   = 11;
    <a id="L48"></a>_DNS_TypePTR   = 12;
    <a id="L49"></a>_DNS_TypeHINFO = 13;
    <a id="L50"></a>_DNS_TypeMINFO = 14;
    <a id="L51"></a>_DNS_TypeMX    = 15;
    <a id="L52"></a>_DNS_TypeTXT   = 16;

    <a id="L54"></a><span class="comment">// valid _DNS_Question.qtype only</span>
    <a id="L55"></a>_DNS_TypeAXFR  = 252;
    <a id="L56"></a>_DNS_TypeMAILB = 253;
    <a id="L57"></a>_DNS_TypeMAILA = 254;
    <a id="L58"></a>_DNS_TypeALL   = 255;

    <a id="L60"></a><span class="comment">// valid _DNS_Question.qclass</span>
    <a id="L61"></a>_DNS_ClassINET   = 1;
    <a id="L62"></a>_DNS_ClassCSNET  = 2;
    <a id="L63"></a>_DNS_ClassCHAOS  = 3;
    <a id="L64"></a>_DNS_ClassHESIOD = 4;
    <a id="L65"></a>_DNS_ClassANY    = 255;

    <a id="L67"></a><span class="comment">// _DNS_Msg.rcode</span>
    <a id="L68"></a>_DNS_RcodeSuccess        = 0;
    <a id="L69"></a>_DNS_RcodeFormatError    = 1;
    <a id="L70"></a>_DNS_RcodeServerFailure  = 2;
    <a id="L71"></a>_DNS_RcodeNameError      = 3;
    <a id="L72"></a>_DNS_RcodeNotImplemented = 4;
    <a id="L73"></a>_DNS_RcodeRefused        = 5;
<a id="L74"></a>)

<a id="L76"></a><span class="comment">// The wire format for the DNS packet header.</span>
<a id="L77"></a>type __DNS_Header struct {
    <a id="L78"></a>Id                                 uint16;
    <a id="L79"></a>Bits                               uint16;
    <a id="L80"></a>Qdcount, Ancount, Nscount, Arcount uint16;
<a id="L81"></a>}

<a id="L83"></a>const (
    <a id="L84"></a><span class="comment">// __DNS_Header.Bits</span>
    <a id="L85"></a>_QR = 1 &lt;&lt; 15; <span class="comment">// query/response (response=1)</span>
    <a id="L86"></a>_AA = 1 &lt;&lt; 10; <span class="comment">// authoritative</span>
    <a id="L87"></a>_TC = 1 &lt;&lt; 9;  <span class="comment">// truncated</span>
    <a id="L88"></a>_RD = 1 &lt;&lt; 8;  <span class="comment">// recursion desired</span>
    <a id="L89"></a>_RA = 1 &lt;&lt; 7;  <span class="comment">// recursion available</span>
<a id="L90"></a>)

<a id="L92"></a><span class="comment">// DNS queries.</span>
<a id="L93"></a>type _DNS_Question struct {
    <a id="L94"></a>Name   string &#34;domain-name&#34;; <span class="comment">// &#34;domain-name&#34; specifies encoding; see packers below</span>
    <a id="L95"></a>Qtype  uint16;
    <a id="L96"></a>Qclass uint16;
<a id="L97"></a>}

<a id="L99"></a><span class="comment">// DNS responses (resource records).</span>
<a id="L100"></a><span class="comment">// There are many types of messages,</span>
<a id="L101"></a><span class="comment">// but they all share the same header.</span>
<a id="L102"></a>type _DNS_RR_Header struct {
    <a id="L103"></a>Name     string &#34;domain-name&#34;;
    <a id="L104"></a>Rrtype   uint16;
    <a id="L105"></a>Class    uint16;
    <a id="L106"></a>Ttl      uint32;
    <a id="L107"></a>Rdlength uint16; <span class="comment">// length of data after header</span>
<a id="L108"></a>}

<a id="L110"></a>func (h *_DNS_RR_Header) Header() *_DNS_RR_Header {
    <a id="L111"></a>return h
<a id="L112"></a>}

<a id="L114"></a>type _DNS_RR interface {
    <a id="L115"></a>Header() *_DNS_RR_Header;
<a id="L116"></a>}


<a id="L119"></a><span class="comment">// Specific DNS RR formats for each query type.</span>

<a id="L121"></a>type _DNS_RR_CNAME struct {
    <a id="L122"></a>Hdr   _DNS_RR_Header;
    <a id="L123"></a>Cname string &#34;domain-name&#34;;
<a id="L124"></a>}

<a id="L126"></a>func (rr *_DNS_RR_CNAME) Header() *_DNS_RR_Header {
    <a id="L127"></a>return &amp;rr.Hdr
<a id="L128"></a>}

<a id="L130"></a>type _DNS_RR_HINFO struct {
    <a id="L131"></a>Hdr _DNS_RR_Header;
    <a id="L132"></a>Cpu string;
    <a id="L133"></a>Os  string;
<a id="L134"></a>}

<a id="L136"></a>func (rr *_DNS_RR_HINFO) Header() *_DNS_RR_Header {
    <a id="L137"></a>return &amp;rr.Hdr
<a id="L138"></a>}

<a id="L140"></a>type _DNS_RR_MB struct {
    <a id="L141"></a>Hdr _DNS_RR_Header;
    <a id="L142"></a>Mb  string &#34;domain-name&#34;;
<a id="L143"></a>}

<a id="L145"></a>func (rr *_DNS_RR_MB) Header() *_DNS_RR_Header {
    <a id="L146"></a>return &amp;rr.Hdr
<a id="L147"></a>}

<a id="L149"></a>type _DNS_RR_MG struct {
    <a id="L150"></a>Hdr _DNS_RR_Header;
    <a id="L151"></a>Mg  string &#34;domain-name&#34;;
<a id="L152"></a>}

<a id="L154"></a>func (rr *_DNS_RR_MG) Header() *_DNS_RR_Header {
    <a id="L155"></a>return &amp;rr.Hdr
<a id="L156"></a>}

<a id="L158"></a>type _DNS_RR_MINFO struct {
    <a id="L159"></a>Hdr   _DNS_RR_Header;
    <a id="L160"></a>Rmail string &#34;domain-name&#34;;
    <a id="L161"></a>Email string &#34;domain-name&#34;;
<a id="L162"></a>}

<a id="L164"></a>func (rr *_DNS_RR_MINFO) Header() *_DNS_RR_Header {
    <a id="L165"></a>return &amp;rr.Hdr
<a id="L166"></a>}

<a id="L168"></a>type _DNS_RR_MR struct {
    <a id="L169"></a>Hdr _DNS_RR_Header;
    <a id="L170"></a>Mr  string &#34;domain-name&#34;;
<a id="L171"></a>}

<a id="L173"></a>func (rr *_DNS_RR_MR) Header() *_DNS_RR_Header {
    <a id="L174"></a>return &amp;rr.Hdr
<a id="L175"></a>}

<a id="L177"></a>type _DNS_RR_MX struct {
    <a id="L178"></a>Hdr  _DNS_RR_Header;
    <a id="L179"></a>Pref uint16;
    <a id="L180"></a>Mx   string &#34;domain-name&#34;;
<a id="L181"></a>}

<a id="L183"></a>func (rr *_DNS_RR_MX) Header() *_DNS_RR_Header {
    <a id="L184"></a>return &amp;rr.Hdr
<a id="L185"></a>}

<a id="L187"></a>type _DNS_RR_NS struct {
    <a id="L188"></a>Hdr _DNS_RR_Header;
    <a id="L189"></a>Ns  string &#34;domain-name&#34;;
<a id="L190"></a>}

<a id="L192"></a>func (rr *_DNS_RR_NS) Header() *_DNS_RR_Header {
    <a id="L193"></a>return &amp;rr.Hdr
<a id="L194"></a>}

<a id="L196"></a>type _DNS_RR_PTR struct {
    <a id="L197"></a>Hdr _DNS_RR_Header;
    <a id="L198"></a>Ptr string &#34;domain-name&#34;;
<a id="L199"></a>}

<a id="L201"></a>func (rr *_DNS_RR_PTR) Header() *_DNS_RR_Header {
    <a id="L202"></a>return &amp;rr.Hdr
<a id="L203"></a>}

<a id="L205"></a>type _DNS_RR_SOA struct {
    <a id="L206"></a>Hdr     _DNS_RR_Header;
    <a id="L207"></a>Ns      string &#34;domain-name&#34;;
    <a id="L208"></a>Mbox    string &#34;domain-name&#34;;
    <a id="L209"></a>Serial  uint32;
    <a id="L210"></a>Refresh uint32;
    <a id="L211"></a>Retry   uint32;
    <a id="L212"></a>Expire  uint32;
    <a id="L213"></a>Minttl  uint32;
<a id="L214"></a>}

<a id="L216"></a>func (rr *_DNS_RR_SOA) Header() *_DNS_RR_Header {
    <a id="L217"></a>return &amp;rr.Hdr
<a id="L218"></a>}

<a id="L220"></a>type _DNS_RR_TXT struct {
    <a id="L221"></a>Hdr _DNS_RR_Header;
    <a id="L222"></a>Txt string; <span class="comment">// not domain name</span>
<a id="L223"></a>}

<a id="L225"></a>func (rr *_DNS_RR_TXT) Header() *_DNS_RR_Header {
    <a id="L226"></a>return &amp;rr.Hdr
<a id="L227"></a>}

<a id="L229"></a>type _DNS_RR_A struct {
    <a id="L230"></a>Hdr _DNS_RR_Header;
    <a id="L231"></a>A   uint32 &#34;ipv4&#34;;
<a id="L232"></a>}

<a id="L234"></a>func (rr *_DNS_RR_A) Header() *_DNS_RR_Header { return &amp;rr.Hdr }


<a id="L237"></a><span class="comment">// Packing and unpacking.</span>
<a id="L238"></a><span class="comment">//</span>
<a id="L239"></a><span class="comment">// All the packers and unpackers take a (msg []byte, off int)</span>
<a id="L240"></a><span class="comment">// and return (off1 int, ok bool).  If they return ok==false, they</span>
<a id="L241"></a><span class="comment">// also return off1==len(msg), so that the next unpacker will</span>
<a id="L242"></a><span class="comment">// also fail.  This lets us avoid checks of ok until the end of a</span>
<a id="L243"></a><span class="comment">// packing sequence.</span>

<a id="L245"></a><span class="comment">// Map of constructors for each RR wire type.</span>
<a id="L246"></a>var rr_mk = map[int]func() _DNS_RR{
    <a id="L247"></a>_DNS_TypeCNAME: func() _DNS_RR { return new(_DNS_RR_CNAME) },
    <a id="L248"></a>_DNS_TypeHINFO: func() _DNS_RR { return new(_DNS_RR_HINFO) },
    <a id="L249"></a>_DNS_TypeMB: func() _DNS_RR { return new(_DNS_RR_MB) },
    <a id="L250"></a>_DNS_TypeMG: func() _DNS_RR { return new(_DNS_RR_MG) },
    <a id="L251"></a>_DNS_TypeMINFO: func() _DNS_RR { return new(_DNS_RR_MINFO) },
    <a id="L252"></a>_DNS_TypeMR: func() _DNS_RR { return new(_DNS_RR_MR) },
    <a id="L253"></a>_DNS_TypeMX: func() _DNS_RR { return new(_DNS_RR_MX) },
    <a id="L254"></a>_DNS_TypeNS: func() _DNS_RR { return new(_DNS_RR_NS) },
    <a id="L255"></a>_DNS_TypePTR: func() _DNS_RR { return new(_DNS_RR_PTR) },
    <a id="L256"></a>_DNS_TypeSOA: func() _DNS_RR { return new(_DNS_RR_SOA) },
    <a id="L257"></a>_DNS_TypeTXT: func() _DNS_RR { return new(_DNS_RR_TXT) },
    <a id="L258"></a>_DNS_TypeA: func() _DNS_RR { return new(_DNS_RR_A) },
<a id="L259"></a>}

<a id="L261"></a><span class="comment">// Pack a domain name s into msg[off:].</span>
<a id="L262"></a><span class="comment">// Domain names are a sequence of counted strings</span>
<a id="L263"></a><span class="comment">// split at the dots.  They end with a zero-length string.</span>
<a id="L264"></a>func packDomainName(s string, msg []byte, off int) (off1 int, ok bool) {
    <a id="L265"></a><span class="comment">// Add trailing dot to canonicalize name.</span>
    <a id="L266"></a>if n := len(s); n == 0 || s[n-1] != &#39;.&#39; {
        <a id="L267"></a>s += &#34;.&#34;
    <a id="L268"></a>}

    <a id="L270"></a><span class="comment">// Each dot ends a segment of the name.</span>
    <a id="L271"></a><span class="comment">// We trade each dot byte for a length byte.</span>
    <a id="L272"></a><span class="comment">// There is also a trailing zero.</span>
    <a id="L273"></a><span class="comment">// Check that we have all the space we need.</span>
    <a id="L274"></a>tot := len(s) + 1;
    <a id="L275"></a>if off+tot &gt; len(msg) {
        <a id="L276"></a>return len(msg), false
    <a id="L277"></a>}

    <a id="L279"></a><span class="comment">// Emit sequence of counted strings, chopping at dots.</span>
    <a id="L280"></a>begin := 0;
    <a id="L281"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L282"></a>if s[i] == &#39;.&#39; {
            <a id="L283"></a>if i-begin &gt;= 1&lt;&lt;6 { <span class="comment">// top two bits of length must be clear</span>
                <a id="L284"></a>return len(msg), false
            <a id="L285"></a>}
            <a id="L286"></a>msg[off] = byte(i - begin);
            <a id="L287"></a>off++;
            <a id="L288"></a>for j := begin; j &lt; i; j++ {
                <a id="L289"></a>msg[off] = s[j];
                <a id="L290"></a>off++;
            <a id="L291"></a>}
            <a id="L292"></a>begin = i + 1;
        <a id="L293"></a>}
    <a id="L294"></a>}
    <a id="L295"></a>msg[off] = 0;
    <a id="L296"></a>off++;
    <a id="L297"></a>return off, true;
<a id="L298"></a>}

<a id="L300"></a><span class="comment">// Unpack a domain name.</span>
<a id="L301"></a><span class="comment">// In addition to the simple sequences of counted strings above,</span>
<a id="L302"></a><span class="comment">// domain names are allowed to refer to strings elsewhere in the</span>
<a id="L303"></a><span class="comment">// packet, to avoid repeating common suffixes when returning</span>
<a id="L304"></a><span class="comment">// many entries in a single domain.  The pointers are marked</span>
<a id="L305"></a><span class="comment">// by a length byte with the top two bits set.  Ignoring those</span>
<a id="L306"></a><span class="comment">// two bits, that byte and the next give a 14 bit offset from msg[0]</span>
<a id="L307"></a><span class="comment">// where we should pick up the trail.</span>
<a id="L308"></a><span class="comment">// Note that if we jump elsewhere in the packet,</span>
<a id="L309"></a><span class="comment">// we return off1 == the offset after the first pointer we found,</span>
<a id="L310"></a><span class="comment">// which is where the next record will start.</span>
<a id="L311"></a><span class="comment">// In theory, the pointers are only allowed to jump backward.</span>
<a id="L312"></a><span class="comment">// We let them jump anywhere and stop jumping after a while.</span>
<a id="L313"></a>func unpackDomainName(msg []byte, off int) (s string, off1 int, ok bool) {
    <a id="L314"></a>s = &#34;&#34;;
    <a id="L315"></a>ptr := 0; <span class="comment">// number of pointers followed</span>
<a id="L316"></a>Loop:
    <a id="L317"></a>for {
        <a id="L318"></a>if off &gt;= len(msg) {
            <a id="L319"></a>return &#34;&#34;, len(msg), false
        <a id="L320"></a>}
        <a id="L321"></a>c := int(msg[off]);
        <a id="L322"></a>off++;
        <a id="L323"></a>switch c &amp; 0xC0 {
        <a id="L324"></a>case 0x00:
            <a id="L325"></a>if c == 0x00 {
                <a id="L326"></a><span class="comment">// end of name</span>
                <a id="L327"></a>break Loop
            <a id="L328"></a>}
            <a id="L329"></a><span class="comment">// literal string</span>
            <a id="L330"></a>if off+c &gt; len(msg) {
                <a id="L331"></a>return &#34;&#34;, len(msg), false
            <a id="L332"></a>}
            <a id="L333"></a>s += string(msg[off:off+c]) + &#34;.&#34;;
            <a id="L334"></a>off += c;
        <a id="L335"></a>case 0xC0:
            <a id="L336"></a><span class="comment">// pointer to somewhere else in msg.</span>
            <a id="L337"></a><span class="comment">// remember location after first ptr,</span>
            <a id="L338"></a><span class="comment">// since that&#39;s how many bytes we consumed.</span>
            <a id="L339"></a><span class="comment">// also, don&#39;t follow too many pointers --</span>
            <a id="L340"></a><span class="comment">// maybe there&#39;s a loop.</span>
            <a id="L341"></a>if off &gt;= len(msg) {
                <a id="L342"></a>return &#34;&#34;, len(msg), false
            <a id="L343"></a>}
            <a id="L344"></a>c1 := msg[off];
            <a id="L345"></a>off++;
            <a id="L346"></a>if ptr == 0 {
                <a id="L347"></a>off1 = off
            <a id="L348"></a>}
            <a id="L349"></a>if ptr++; ptr &gt; 10 {
                <a id="L350"></a>return &#34;&#34;, len(msg), false
            <a id="L351"></a>}
            <a id="L352"></a>off = (c^0xC0)&lt;&lt;8 | int(c1);
        <a id="L353"></a>default:
            <a id="L354"></a><span class="comment">// 0x80 and 0x40 are reserved</span>
            <a id="L355"></a>return &#34;&#34;, len(msg), false
        <a id="L356"></a>}
    <a id="L357"></a>}
    <a id="L358"></a>if ptr == 0 {
        <a id="L359"></a>off1 = off
    <a id="L360"></a>}
    <a id="L361"></a>return s, off1, true;
<a id="L362"></a>}

<a id="L364"></a><span class="comment">// TODO(rsc): Move into generic library?</span>
<a id="L365"></a><span class="comment">// Pack a reflect.StructValue into msg.  Struct members can only be uint16, uint32, string,</span>
<a id="L366"></a><span class="comment">// and other (often anonymous) structs.</span>
<a id="L367"></a>func packStructValue(val *reflect.StructValue, msg []byte, off int) (off1 int, ok bool) {
    <a id="L368"></a>for i := 0; i &lt; val.NumField(); i++ {
        <a id="L369"></a>f := val.Type().(*reflect.StructType).Field(i);
        <a id="L370"></a>switch fv := val.Field(i).(type) {
        <a id="L371"></a>default:
            <a id="L372"></a>fmt.Fprintf(os.Stderr, &#34;net: dns: unknown packing type %v&#34;, f.Type);
            <a id="L373"></a>return len(msg), false;
        <a id="L374"></a>case *reflect.StructValue:
            <a id="L375"></a>off, ok = packStructValue(fv, msg, off)
        <a id="L376"></a>case *reflect.Uint16Value:
            <a id="L377"></a>i := fv.Get();
            <a id="L378"></a>if off+2 &gt; len(msg) {
                <a id="L379"></a>return len(msg), false
            <a id="L380"></a>}
            <a id="L381"></a>msg[off] = byte(i &gt;&gt; 8);
            <a id="L382"></a>msg[off+1] = byte(i);
            <a id="L383"></a>off += 2;
        <a id="L384"></a>case *reflect.Uint32Value:
            <a id="L385"></a>i := fv.Get();
            <a id="L386"></a>if off+4 &gt; len(msg) {
                <a id="L387"></a>return len(msg), false
            <a id="L388"></a>}
            <a id="L389"></a>msg[off] = byte(i &gt;&gt; 24);
            <a id="L390"></a>msg[off+1] = byte(i &gt;&gt; 16);
            <a id="L391"></a>msg[off+2] = byte(i &gt;&gt; 8);
            <a id="L392"></a>msg[off+4] = byte(i);
            <a id="L393"></a>off += 4;
        <a id="L394"></a>case *reflect.StringValue:
            <a id="L395"></a><span class="comment">// There are multiple string encodings.</span>
            <a id="L396"></a><span class="comment">// The tag distinguishes ordinary strings from domain names.</span>
            <a id="L397"></a>s := fv.Get();
            <a id="L398"></a>switch f.Tag {
            <a id="L399"></a>default:
                <a id="L400"></a>fmt.Fprintf(os.Stderr, &#34;net: dns: unknown string tag %v&#34;, f.Tag);
                <a id="L401"></a>return len(msg), false;
            <a id="L402"></a>case &#34;domain-name&#34;:
                <a id="L403"></a>off, ok = packDomainName(s, msg, off);
                <a id="L404"></a>if !ok {
                    <a id="L405"></a>return len(msg), false
                <a id="L406"></a>}
            <a id="L407"></a>case &#34;&#34;:
                <a id="L408"></a><span class="comment">// Counted string: 1 byte length.</span>
                <a id="L409"></a>if len(s) &gt; 255 || off+1+len(s) &gt; len(msg) {
                    <a id="L410"></a>return len(msg), false
                <a id="L411"></a>}
                <a id="L412"></a>msg[off] = byte(len(s));
                <a id="L413"></a>off++;
                <a id="L414"></a>for i := 0; i &lt; len(s); i++ {
                    <a id="L415"></a>msg[off+i] = s[i]
                <a id="L416"></a>}
                <a id="L417"></a>off += len(s);
            <a id="L418"></a>}
        <a id="L419"></a>}
    <a id="L420"></a>}
    <a id="L421"></a>return off, true;
<a id="L422"></a>}

<a id="L424"></a>func structValue(any interface{}) *reflect.StructValue {
    <a id="L425"></a>return reflect.NewValue(any).(*reflect.PtrValue).Elem().(*reflect.StructValue)
<a id="L426"></a>}

<a id="L428"></a>func packStruct(any interface{}, msg []byte, off int) (off1 int, ok bool) {
    <a id="L429"></a>off, ok = packStructValue(structValue(any), msg, off);
    <a id="L430"></a>return off, ok;
<a id="L431"></a>}

<a id="L433"></a><span class="comment">// TODO(rsc): Move into generic library?</span>
<a id="L434"></a><span class="comment">// Unpack a reflect.StructValue from msg.</span>
<a id="L435"></a><span class="comment">// Same restrictions as packStructValue.</span>
<a id="L436"></a>func unpackStructValue(val *reflect.StructValue, msg []byte, off int) (off1 int, ok bool) {
    <a id="L437"></a>for i := 0; i &lt; val.NumField(); i++ {
        <a id="L438"></a>f := val.Type().(*reflect.StructType).Field(i);
        <a id="L439"></a>switch fv := val.Field(i).(type) {
        <a id="L440"></a>default:
            <a id="L441"></a>fmt.Fprintf(os.Stderr, &#34;net: dns: unknown packing type %v&#34;, f.Type);
            <a id="L442"></a>return len(msg), false;
        <a id="L443"></a>case *reflect.StructValue:
            <a id="L444"></a>off, ok = unpackStructValue(fv, msg, off)
        <a id="L445"></a>case *reflect.Uint16Value:
            <a id="L446"></a>if off+2 &gt; len(msg) {
                <a id="L447"></a>return len(msg), false
            <a id="L448"></a>}
            <a id="L449"></a>i := uint16(msg[off])&lt;&lt;8 | uint16(msg[off+1]);
            <a id="L450"></a>fv.Set(i);
            <a id="L451"></a>off += 2;
        <a id="L452"></a>case *reflect.Uint32Value:
            <a id="L453"></a>if off+4 &gt; len(msg) {
                <a id="L454"></a>return len(msg), false
            <a id="L455"></a>}
            <a id="L456"></a>i := uint32(msg[off])&lt;&lt;24 | uint32(msg[off+1])&lt;&lt;16 | uint32(msg[off+2])&lt;&lt;8 | uint32(msg[off+3]);
            <a id="L457"></a>fv.Set(i);
            <a id="L458"></a>off += 4;
        <a id="L459"></a>case *reflect.StringValue:
            <a id="L460"></a>var s string;
            <a id="L461"></a>switch f.Tag {
            <a id="L462"></a>default:
                <a id="L463"></a>fmt.Fprintf(os.Stderr, &#34;net: dns: unknown string tag %v&#34;, f.Tag);
                <a id="L464"></a>return len(msg), false;
            <a id="L465"></a>case &#34;domain-name&#34;:
                <a id="L466"></a>s, off, ok = unpackDomainName(msg, off);
                <a id="L467"></a>if !ok {
                    <a id="L468"></a>return len(msg), false
                <a id="L469"></a>}
            <a id="L470"></a>case &#34;&#34;:
                <a id="L471"></a>if off &gt;= len(msg) || off+1+int(msg[off]) &gt; len(msg) {
                    <a id="L472"></a>return len(msg), false
                <a id="L473"></a>}
                <a id="L474"></a>n := int(msg[off]);
                <a id="L475"></a>off++;
                <a id="L476"></a>b := make([]byte, n);
                <a id="L477"></a>for i := 0; i &lt; n; i++ {
                    <a id="L478"></a>b[i] = msg[off+i]
                <a id="L479"></a>}
                <a id="L480"></a>off += n;
                <a id="L481"></a>s = string(b);
            <a id="L482"></a>}
            <a id="L483"></a>fv.Set(s);
        <a id="L484"></a>}
    <a id="L485"></a>}
    <a id="L486"></a>return off, true;
<a id="L487"></a>}

<a id="L489"></a>func unpackStruct(any interface{}, msg []byte, off int) (off1 int, ok bool) {
    <a id="L490"></a>off, ok = unpackStructValue(structValue(any), msg, off);
    <a id="L491"></a>return off, ok;
<a id="L492"></a>}

<a id="L494"></a><span class="comment">// Generic struct printer.</span>
<a id="L495"></a><span class="comment">// Doesn&#39;t care about the string tag &#34;domain-name&#34;,</span>
<a id="L496"></a><span class="comment">// but does look for an &#34;ipv4&#34; tag on uint32 variables,</span>
<a id="L497"></a><span class="comment">// printing them as IP addresses.</span>
<a id="L498"></a>func printStructValue(val *reflect.StructValue) string {
    <a id="L499"></a>s := &#34;{&#34;;
    <a id="L500"></a>for i := 0; i &lt; val.NumField(); i++ {
        <a id="L501"></a>if i &gt; 0 {
            <a id="L502"></a>s += &#34;, &#34;
        <a id="L503"></a>}
        <a id="L504"></a>f := val.Type().(*reflect.StructType).Field(i);
        <a id="L505"></a>if !f.Anonymous {
            <a id="L506"></a>s += f.Name + &#34;=&#34;
        <a id="L507"></a>}
        <a id="L508"></a>fval := val.Field(i);
        <a id="L509"></a>if fv, ok := fval.(*reflect.StructValue); ok {
            <a id="L510"></a>s += printStructValue(fv)
        <a id="L511"></a>} else if fv, ok := fval.(*reflect.Uint32Value); ok &amp;&amp; f.Tag == &#34;ipv4&#34; {
            <a id="L512"></a>i := fv.Get();
            <a id="L513"></a>s += IPv4(byte(i&gt;&gt;24), byte(i&gt;&gt;16), byte(i&gt;&gt;8), byte(i)).String();
        <a id="L514"></a>} else {
            <a id="L515"></a>s += fmt.Sprint(fval.Interface())
        <a id="L516"></a>}
    <a id="L517"></a>}
    <a id="L518"></a>s += &#34;}&#34;;
    <a id="L519"></a>return s;
<a id="L520"></a>}

<a id="L522"></a>func printStruct(any interface{}) string { return printStructValue(structValue(any)) }

<a id="L524"></a><span class="comment">// Resource record packer.</span>
<a id="L525"></a>func packRR(rr _DNS_RR, msg []byte, off int) (off2 int, ok bool) {
    <a id="L526"></a>var off1 int;
    <a id="L527"></a><span class="comment">// pack twice, once to find end of header</span>
    <a id="L528"></a><span class="comment">// and again to find end of packet.</span>
    <a id="L529"></a><span class="comment">// a bit inefficient but this doesn&#39;t need to be fast.</span>
    <a id="L530"></a><span class="comment">// off1 is end of header</span>
    <a id="L531"></a><span class="comment">// off2 is end of rr</span>
    <a id="L532"></a>off1, ok = packStruct(rr.Header(), msg, off);
    <a id="L533"></a>off2, ok = packStruct(rr, msg, off);
    <a id="L534"></a>if !ok {
        <a id="L535"></a>return len(msg), false
    <a id="L536"></a>}
    <a id="L537"></a><span class="comment">// pack a third time; redo header with correct data length</span>
    <a id="L538"></a>rr.Header().Rdlength = uint16(off2 - off1);
    <a id="L539"></a>packStruct(rr.Header(), msg, off);
    <a id="L540"></a>return off2, true;
<a id="L541"></a>}

<a id="L543"></a><span class="comment">// Resource record unpacker.</span>
<a id="L544"></a>func unpackRR(msg []byte, off int) (rr _DNS_RR, off1 int, ok bool) {
    <a id="L545"></a><span class="comment">// unpack just the header, to find the rr type and length</span>
    <a id="L546"></a>var h _DNS_RR_Header;
    <a id="L547"></a>off0 := off;
    <a id="L548"></a>if off, ok = unpackStruct(&amp;h, msg, off); !ok {
        <a id="L549"></a>return nil, len(msg), false
    <a id="L550"></a>}
    <a id="L551"></a>end := off + int(h.Rdlength);

    <a id="L553"></a><span class="comment">// make an rr of that type and re-unpack.</span>
    <a id="L554"></a><span class="comment">// again inefficient but doesn&#39;t need to be fast.</span>
    <a id="L555"></a>mk, known := rr_mk[int(h.Rrtype)];
    <a id="L556"></a>if !known {
        <a id="L557"></a>return &amp;h, end, true
    <a id="L558"></a>}
    <a id="L559"></a>rr = mk();
    <a id="L560"></a>off, ok = unpackStruct(rr, msg, off0);
    <a id="L561"></a>if off != end {
        <a id="L562"></a>return &amp;h, end, true
    <a id="L563"></a>}
    <a id="L564"></a>return rr, off, ok;
<a id="L565"></a>}

<a id="L567"></a><span class="comment">// Usable representation of a DNS packet.</span>

<a id="L569"></a><span class="comment">// A manually-unpacked version of (id, bits).</span>
<a id="L570"></a><span class="comment">// This is in its own struct for easy printing.</span>
<a id="L571"></a>type __DNS_Msg_Top struct {
    <a id="L572"></a>id                  uint16;
    <a id="L573"></a>response            bool;
    <a id="L574"></a>opcode              int;
    <a id="L575"></a>authoritative       bool;
    <a id="L576"></a>truncated           bool;
    <a id="L577"></a>recursion_desired   bool;
    <a id="L578"></a>recursion_available bool;
    <a id="L579"></a>rcode               int;
<a id="L580"></a>}

<a id="L582"></a>type _DNS_Msg struct {
    <a id="L583"></a>__DNS_Msg_Top;
    <a id="L584"></a>question []_DNS_Question;
    <a id="L585"></a>answer   []_DNS_RR;
    <a id="L586"></a>ns       []_DNS_RR;
    <a id="L587"></a>extra    []_DNS_RR;
<a id="L588"></a>}


<a id="L591"></a>func (dns *_DNS_Msg) Pack() (msg []byte, ok bool) {
    <a id="L592"></a>var dh __DNS_Header;

    <a id="L594"></a><span class="comment">// Convert convenient _DNS_Msg into wire-like __DNS_Header.</span>
    <a id="L595"></a>dh.Id = dns.id;
    <a id="L596"></a>dh.Bits = uint16(dns.opcode)&lt;&lt;11 | uint16(dns.rcode);
    <a id="L597"></a>if dns.recursion_available {
        <a id="L598"></a>dh.Bits |= _RA
    <a id="L599"></a>}
    <a id="L600"></a>if dns.recursion_desired {
        <a id="L601"></a>dh.Bits |= _RD
    <a id="L602"></a>}
    <a id="L603"></a>if dns.truncated {
        <a id="L604"></a>dh.Bits |= _TC
    <a id="L605"></a>}
    <a id="L606"></a>if dns.authoritative {
        <a id="L607"></a>dh.Bits |= _AA
    <a id="L608"></a>}
    <a id="L609"></a>if dns.response {
        <a id="L610"></a>dh.Bits |= _QR
    <a id="L611"></a>}

    <a id="L613"></a><span class="comment">// Prepare variable sized arrays.</span>
    <a id="L614"></a>question := dns.question;
    <a id="L615"></a>answer := dns.answer;
    <a id="L616"></a>ns := dns.ns;
    <a id="L617"></a>extra := dns.extra;

    <a id="L619"></a>dh.Qdcount = uint16(len(question));
    <a id="L620"></a>dh.Ancount = uint16(len(answer));
    <a id="L621"></a>dh.Nscount = uint16(len(ns));
    <a id="L622"></a>dh.Arcount = uint16(len(extra));

    <a id="L624"></a><span class="comment">// Could work harder to calculate message size,</span>
    <a id="L625"></a><span class="comment">// but this is far more than we need and not</span>
    <a id="L626"></a><span class="comment">// big enough to hurt the allocator.</span>
    <a id="L627"></a>msg = make([]byte, 2000);

    <a id="L629"></a><span class="comment">// Pack it in: header and then the pieces.</span>
    <a id="L630"></a>off := 0;
    <a id="L631"></a>off, ok = packStruct(&amp;dh, msg, off);
    <a id="L632"></a>for i := 0; i &lt; len(question); i++ {
        <a id="L633"></a>off, ok = packStruct(&amp;question[i], msg, off)
    <a id="L634"></a>}
    <a id="L635"></a>for i := 0; i &lt; len(answer); i++ {
        <a id="L636"></a>off, ok = packStruct(answer[i], msg, off)
    <a id="L637"></a>}
    <a id="L638"></a>for i := 0; i &lt; len(ns); i++ {
        <a id="L639"></a>off, ok = packStruct(ns[i], msg, off)
    <a id="L640"></a>}
    <a id="L641"></a>for i := 0; i &lt; len(extra); i++ {
        <a id="L642"></a>off, ok = packStruct(extra[i], msg, off)
    <a id="L643"></a>}
    <a id="L644"></a>if !ok {
        <a id="L645"></a>return nil, false
    <a id="L646"></a>}
    <a id="L647"></a>return msg[0:off], true;
<a id="L648"></a>}

<a id="L650"></a>func (dns *_DNS_Msg) Unpack(msg []byte) bool {
    <a id="L651"></a><span class="comment">// Header.</span>
    <a id="L652"></a>var dh __DNS_Header;
    <a id="L653"></a>off := 0;
    <a id="L654"></a>var ok bool;
    <a id="L655"></a>if off, ok = unpackStruct(&amp;dh, msg, off); !ok {
        <a id="L656"></a>return false
    <a id="L657"></a>}
    <a id="L658"></a>dns.id = dh.Id;
    <a id="L659"></a>dns.response = (dh.Bits &amp; _QR) != 0;
    <a id="L660"></a>dns.opcode = int(dh.Bits&gt;&gt;11) &amp; 0xF;
    <a id="L661"></a>dns.authoritative = (dh.Bits &amp; _AA) != 0;
    <a id="L662"></a>dns.truncated = (dh.Bits &amp; _TC) != 0;
    <a id="L663"></a>dns.recursion_desired = (dh.Bits &amp; _RD) != 0;
    <a id="L664"></a>dns.recursion_available = (dh.Bits &amp; _RA) != 0;
    <a id="L665"></a>dns.rcode = int(dh.Bits &amp; 0xF);

    <a id="L667"></a><span class="comment">// Arrays.</span>
    <a id="L668"></a>dns.question = make([]_DNS_Question, dh.Qdcount);
    <a id="L669"></a>dns.answer = make([]_DNS_RR, dh.Ancount);
    <a id="L670"></a>dns.ns = make([]_DNS_RR, dh.Nscount);
    <a id="L671"></a>dns.extra = make([]_DNS_RR, dh.Arcount);

    <a id="L673"></a>for i := 0; i &lt; len(dns.question); i++ {
        <a id="L674"></a>off, ok = unpackStruct(&amp;dns.question[i], msg, off)
    <a id="L675"></a>}
    <a id="L676"></a>for i := 0; i &lt; len(dns.answer); i++ {
        <a id="L677"></a>dns.answer[i], off, ok = unpackRR(msg, off)
    <a id="L678"></a>}
    <a id="L679"></a>for i := 0; i &lt; len(dns.ns); i++ {
        <a id="L680"></a>dns.ns[i], off, ok = unpackRR(msg, off)
    <a id="L681"></a>}
    <a id="L682"></a>for i := 0; i &lt; len(dns.extra); i++ {
        <a id="L683"></a>dns.extra[i], off, ok = unpackRR(msg, off)
    <a id="L684"></a>}
    <a id="L685"></a>if !ok {
        <a id="L686"></a>return false
    <a id="L687"></a>}
    <a id="L688"></a><span class="comment">//	if off != len(msg) {</span>
    <a id="L689"></a><span class="comment">//		println(&#34;extra bytes in dns packet&#34;, off, &#34;&lt;&#34;, len(msg));</span>
    <a id="L690"></a><span class="comment">//	}</span>
    <a id="L691"></a>return true;
<a id="L692"></a>}

<a id="L694"></a>func (dns *_DNS_Msg) String() string {
    <a id="L695"></a>s := &#34;DNS: &#34; + printStruct(&amp;dns.__DNS_Msg_Top) + &#34;\n&#34;;
    <a id="L696"></a>if len(dns.question) &gt; 0 {
        <a id="L697"></a>s += &#34;-- Questions\n&#34;;
        <a id="L698"></a>for i := 0; i &lt; len(dns.question); i++ {
            <a id="L699"></a>s += printStruct(&amp;dns.question[i]) + &#34;\n&#34;
        <a id="L700"></a>}
    <a id="L701"></a>}
    <a id="L702"></a>if len(dns.answer) &gt; 0 {
        <a id="L703"></a>s += &#34;-- Answers\n&#34;;
        <a id="L704"></a>for i := 0; i &lt; len(dns.answer); i++ {
            <a id="L705"></a>s += printStruct(dns.answer[i]) + &#34;\n&#34;
        <a id="L706"></a>}
    <a id="L707"></a>}
    <a id="L708"></a>if len(dns.ns) &gt; 0 {
        <a id="L709"></a>s += &#34;-- Name servers\n&#34;;
        <a id="L710"></a>for i := 0; i &lt; len(dns.ns); i++ {
            <a id="L711"></a>s += printStruct(dns.ns[i]) + &#34;\n&#34;
        <a id="L712"></a>}
    <a id="L713"></a>}
    <a id="L714"></a>if len(dns.extra) &gt; 0 {
        <a id="L715"></a>s += &#34;-- Extra\n&#34;;
        <a id="L716"></a>for i := 0; i &lt; len(dns.extra); i++ {
            <a id="L717"></a>s += printStruct(dns.extra[i]) + &#34;\n&#34;
        <a id="L718"></a>}
    <a id="L719"></a>}
    <a id="L720"></a>return s;
<a id="L721"></a>}
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
