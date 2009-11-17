<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/dnsclient.go</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/net/dnsclient.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// DNS client: see RFC 1035.</span>
<a id="L6"></a><span class="comment">// Has to be linked into package net for Dial.</span>

<a id="L8"></a><span class="comment">// TODO(rsc):</span>
<a id="L9"></a><span class="comment">//	Check periodically whether /etc/resolv.conf has changed.</span>
<a id="L10"></a><span class="comment">//	Could potentially handle many outstanding lookups faster.</span>
<a id="L11"></a><span class="comment">//	Could have a small cache.</span>
<a id="L12"></a><span class="comment">//	Random UDP source port (net.Dial should do that for us).</span>
<a id="L13"></a><span class="comment">//	Random request IDs.</span>

<a id="L15"></a>package net

<a id="L17"></a>import (
    <a id="L18"></a>&#34;once&#34;;
    <a id="L19"></a>&#34;os&#34;;
<a id="L20"></a>)

<a id="L22"></a><span class="comment">// DNSError represents a DNS lookup error.</span>
<a id="L23"></a>type DNSError struct {
    <a id="L24"></a>Error  string; <span class="comment">// description of the error</span>
    <a id="L25"></a>Name   string; <span class="comment">// name looked for</span>
    <a id="L26"></a>Server string; <span class="comment">// server used</span>
<a id="L27"></a>}

<a id="L29"></a>func (e *DNSError) String() string {
    <a id="L30"></a>s := &#34;lookup &#34; + e.Name;
    <a id="L31"></a>if e.Server != &#34;&#34; {
        <a id="L32"></a>s += &#34; on &#34; + e.Server
    <a id="L33"></a>}
    <a id="L34"></a>s += &#34;: &#34; + e.Error;
    <a id="L35"></a>return s;
<a id="L36"></a>}

<a id="L38"></a>const noSuchHost = &#34;no such host&#34;

<a id="L40"></a><span class="comment">// Send a request on the connection and hope for a reply.</span>
<a id="L41"></a><span class="comment">// Up to cfg.attempts attempts.</span>
<a id="L42"></a>func _Exchange(cfg *_DNS_Config, c Conn, name string) (m *_DNS_Msg, err os.Error) {
    <a id="L43"></a>if len(name) &gt;= 256 {
        <a id="L44"></a>return nil, &amp;DNSError{&#34;name too long&#34;, name, &#34;&#34;}
    <a id="L45"></a>}
    <a id="L46"></a>out := new(_DNS_Msg);
    <a id="L47"></a>out.id = 0x1234;
    <a id="L48"></a>out.question = []_DNS_Question{
        <a id="L49"></a>_DNS_Question{name, _DNS_TypeA, _DNS_ClassINET},
    <a id="L50"></a>};
    <a id="L51"></a>out.recursion_desired = true;
    <a id="L52"></a>msg, ok := out.Pack();
    <a id="L53"></a>if !ok {
        <a id="L54"></a>return nil, &amp;DNSError{&#34;internal error - cannot pack message&#34;, name, &#34;&#34;}
    <a id="L55"></a>}

    <a id="L57"></a>for attempt := 0; attempt &lt; cfg.attempts; attempt++ {
        <a id="L58"></a>n, err := c.Write(msg);
        <a id="L59"></a>if err != nil {
            <a id="L60"></a>return nil, err
        <a id="L61"></a>}

        <a id="L63"></a>c.SetReadTimeout(1e9); <span class="comment">// nanoseconds</span>

        <a id="L65"></a>buf := make([]byte, 2000); <span class="comment">// More than enough.</span>
        <a id="L66"></a>n, err = c.Read(buf);
        <a id="L67"></a>if isEAGAIN(err) {
            <a id="L68"></a>err = nil;
            <a id="L69"></a>continue;
        <a id="L70"></a>}
        <a id="L71"></a>if err != nil {
            <a id="L72"></a>return nil, err
        <a id="L73"></a>}
        <a id="L74"></a>buf = buf[0:n];
        <a id="L75"></a>in := new(_DNS_Msg);
        <a id="L76"></a>if !in.Unpack(buf) || in.id != out.id {
            <a id="L77"></a>continue
        <a id="L78"></a>}
        <a id="L79"></a>return in, nil;
    <a id="L80"></a>}
    <a id="L81"></a>var server string;
    <a id="L82"></a>if a := c.RemoteAddr(); a != nil {
        <a id="L83"></a>server = a.String()
    <a id="L84"></a>}
    <a id="L85"></a>return nil, &amp;DNSError{&#34;no answer from server&#34;, name, server};
<a id="L86"></a>}


<a id="L89"></a><span class="comment">// Find answer for name in dns message.</span>
<a id="L90"></a><span class="comment">// On return, if err == nil, addrs != nil.</span>
<a id="L91"></a>func answer(name, server string, dns *_DNS_Msg) (addrs []string, err *DNSError) {
    <a id="L92"></a>addrs = make([]string, 0, len(dns.answer));

    <a id="L94"></a>if dns.rcode == _DNS_RcodeNameError &amp;&amp; dns.recursion_available {
        <a id="L95"></a>return nil, &amp;DNSError{noSuchHost, name, &#34;&#34;}
    <a id="L96"></a>}
    <a id="L97"></a>if dns.rcode != _DNS_RcodeSuccess {
        <a id="L98"></a><span class="comment">// None of the error codes make sense</span>
        <a id="L99"></a><span class="comment">// for the query we sent.  If we didn&#39;t get</span>
        <a id="L100"></a><span class="comment">// a name error and we didn&#39;t get success,</span>
        <a id="L101"></a><span class="comment">// the server is behaving incorrectly.</span>
        <a id="L102"></a>return nil, &amp;DNSError{&#34;server misbehaving&#34;, name, server}
    <a id="L103"></a>}

    <a id="L105"></a><span class="comment">// Look for the name.</span>
    <a id="L106"></a><span class="comment">// Presotto says it&#39;s okay to assume that servers listed in</span>
    <a id="L107"></a><span class="comment">// /etc/resolv.conf are recursive resolvers.</span>
    <a id="L108"></a><span class="comment">// We asked for recursion, so it should have included</span>
    <a id="L109"></a><span class="comment">// all the answers we need in this one packet.</span>
<a id="L110"></a>Cname:
    <a id="L111"></a>for cnameloop := 0; cnameloop &lt; 10; cnameloop++ {
        <a id="L112"></a>addrs = addrs[0:0];
        <a id="L113"></a>for i := 0; i &lt; len(dns.answer); i++ {
            <a id="L114"></a>rr := dns.answer[i];
            <a id="L115"></a>h := rr.Header();
            <a id="L116"></a>if h.Class == _DNS_ClassINET &amp;&amp; h.Name == name {
                <a id="L117"></a>switch h.Rrtype {
                <a id="L118"></a>case _DNS_TypeA:
                    <a id="L119"></a>n := len(addrs);
                    <a id="L120"></a>a := rr.(*_DNS_RR_A).A;
                    <a id="L121"></a>addrs = addrs[0 : n+1];
                    <a id="L122"></a>addrs[n] = IPv4(byte(a&gt;&gt;24), byte(a&gt;&gt;16), byte(a&gt;&gt;8), byte(a)).String();
                <a id="L123"></a>case _DNS_TypeCNAME:
                    <a id="L124"></a><span class="comment">// redirect to cname</span>
                    <a id="L125"></a>name = rr.(*_DNS_RR_CNAME).Cname;
                    <a id="L126"></a>continue Cname;
                <a id="L127"></a>}
            <a id="L128"></a>}
        <a id="L129"></a>}
        <a id="L130"></a>if len(addrs) == 0 {
            <a id="L131"></a>return nil, &amp;DNSError{noSuchHost, name, server}
        <a id="L132"></a>}
        <a id="L133"></a>return addrs, nil;
    <a id="L134"></a>}

    <a id="L136"></a>return nil, &amp;DNSError{&#34;too many redirects&#34;, name, server};
<a id="L137"></a>}

<a id="L139"></a><span class="comment">// Do a lookup for a single name, which must be rooted</span>
<a id="L140"></a><span class="comment">// (otherwise answer will not find the answers).</span>
<a id="L141"></a>func tryOneName(cfg *_DNS_Config, name string) (addrs []string, err os.Error) {
    <a id="L142"></a>if len(cfg.servers) == 0 {
        <a id="L143"></a>return nil, &amp;DNSError{&#34;no DNS servers&#34;, name, &#34;&#34;}
    <a id="L144"></a>}
    <a id="L145"></a>for i := 0; i &lt; len(cfg.servers); i++ {
        <a id="L146"></a><span class="comment">// Calling Dial here is scary -- we have to be sure</span>
        <a id="L147"></a><span class="comment">// not to dial a name that will require a DNS lookup,</span>
        <a id="L148"></a><span class="comment">// or Dial will call back here to translate it.</span>
        <a id="L149"></a><span class="comment">// The DNS config parser has already checked that</span>
        <a id="L150"></a><span class="comment">// all the cfg.servers[i] are IP addresses, which</span>
        <a id="L151"></a><span class="comment">// Dial will use without a DNS lookup.</span>
        <a id="L152"></a>server := cfg.servers[i] + &#34;:53&#34;;
        <a id="L153"></a>c, cerr := Dial(&#34;udp&#34;, &#34;&#34;, server);
        <a id="L154"></a>if cerr != nil {
            <a id="L155"></a>err = cerr;
            <a id="L156"></a>continue;
        <a id="L157"></a>}
        <a id="L158"></a>msg, merr := _Exchange(cfg, c, name);
        <a id="L159"></a>c.Close();
        <a id="L160"></a>if merr != nil {
            <a id="L161"></a>err = merr;
            <a id="L162"></a>continue;
        <a id="L163"></a>}
        <a id="L164"></a>var dnserr *DNSError;
        <a id="L165"></a>addrs, dnserr = answer(name, server, msg);
        <a id="L166"></a>if dnserr != nil {
            <a id="L167"></a>err = dnserr
        <a id="L168"></a>} else {
            <a id="L169"></a>err = nil <span class="comment">// nil os.Error, not nil *DNSError</span>
        <a id="L170"></a>}
        <a id="L171"></a>if dnserr == nil || dnserr.Error == noSuchHost {
            <a id="L172"></a>break
        <a id="L173"></a>}
    <a id="L174"></a>}
    <a id="L175"></a>return;
<a id="L176"></a>}

<a id="L178"></a>var cfg *_DNS_Config
<a id="L179"></a>var dnserr os.Error

<a id="L181"></a>func loadConfig() { cfg, dnserr = _DNS_ReadConfig() }

<a id="L183"></a>func isDomainName(s string) bool {
    <a id="L184"></a><span class="comment">// Requirements on DNS name:</span>
    <a id="L185"></a><span class="comment">//	* must not be empty.</span>
    <a id="L186"></a><span class="comment">//	* must be alphanumeric plus - and .</span>
    <a id="L187"></a><span class="comment">//	* each of the dot-separated elements must begin</span>
    <a id="L188"></a><span class="comment">//	  and end with a letter or digit.</span>
    <a id="L189"></a><span class="comment">//	  RFC 1035 required the element to begin with a letter,</span>
    <a id="L190"></a><span class="comment">//	  but RFC 3696 says this has been relaxed to allow digits too.</span>
    <a id="L191"></a><span class="comment">//	  still, there must be a letter somewhere in the entire name.</span>
    <a id="L192"></a>if len(s) == 0 {
        <a id="L193"></a>return false
    <a id="L194"></a>}
    <a id="L195"></a>if s[len(s)-1] != &#39;.&#39; { <span class="comment">// simplify checking loop: make name end in dot</span>
        <a id="L196"></a>s += &#34;.&#34;
    <a id="L197"></a>}

    <a id="L199"></a>last := byte(&#39;.&#39;);
    <a id="L200"></a>ok := false; <span class="comment">// ok once we&#39;ve seen a letter</span>
    <a id="L201"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L202"></a>c := s[i];
        <a id="L203"></a>switch {
        <a id="L204"></a>default:
            <a id="L205"></a>return false
        <a id="L206"></a>case &#39;a&#39; &lt;= c &amp;&amp; c &lt;= &#39;z&#39; || &#39;A&#39; &lt;= c &amp;&amp; c &lt;= &#39;Z&#39;:
            <a id="L207"></a>ok = true
        <a id="L208"></a>case &#39;0&#39; &lt;= c &amp;&amp; c &lt;= &#39;9&#39;:
            <a id="L209"></a><span class="comment">// fine</span>
        <a id="L210"></a>case c == &#39;-&#39;:
            <a id="L211"></a><span class="comment">// byte before dash cannot be dot</span>
            <a id="L212"></a>if last == &#39;.&#39; {
                <a id="L213"></a>return false
            <a id="L214"></a>}
        <a id="L215"></a>case c == &#39;.&#39;:
            <a id="L216"></a><span class="comment">// byte before dot cannot be dot, dash</span>
            <a id="L217"></a>if last == &#39;.&#39; || last == &#39;-&#39; {
                <a id="L218"></a>return false
            <a id="L219"></a>}
        <a id="L220"></a>}
        <a id="L221"></a>last = c;
    <a id="L222"></a>}

    <a id="L224"></a>return ok;
<a id="L225"></a>}

<a id="L227"></a><span class="comment">// LookupHost looks up the host name using the local DNS resolver.</span>
<a id="L228"></a><span class="comment">// It returns the canonical name for the host and an array of that</span>
<a id="L229"></a><span class="comment">// host&#39;s addresses.</span>
<a id="L230"></a>func LookupHost(name string) (cname string, addrs []string, err os.Error) {
    <a id="L231"></a>if !isDomainName(name) {
        <a id="L232"></a>return name, nil, &amp;DNSError{&#34;invalid domain name&#34;, name, &#34;&#34;}
    <a id="L233"></a>}
    <a id="L234"></a>once.Do(loadConfig);
    <a id="L235"></a>if dnserr != nil || cfg == nil {
        <a id="L236"></a>err = dnserr;
        <a id="L237"></a>return;
    <a id="L238"></a>}

    <a id="L240"></a><span class="comment">// If name is rooted (trailing dot) or has enough dots,</span>
    <a id="L241"></a><span class="comment">// try it by itself first.</span>
    <a id="L242"></a>rooted := len(name) &gt; 0 &amp;&amp; name[len(name)-1] == &#39;.&#39;;
    <a id="L243"></a>if rooted || count(name, &#39;.&#39;) &gt;= cfg.ndots {
        <a id="L244"></a>rname := name;
        <a id="L245"></a>if !rooted {
            <a id="L246"></a>rname += &#34;.&#34;
        <a id="L247"></a>}
        <a id="L248"></a><span class="comment">// Can try as ordinary name.</span>
        <a id="L249"></a>addrs, err = tryOneName(cfg, rname);
        <a id="L250"></a>if err == nil {
            <a id="L251"></a>cname = rname;
            <a id="L252"></a>return;
        <a id="L253"></a>}
    <a id="L254"></a>}
    <a id="L255"></a>if rooted {
        <a id="L256"></a>return
    <a id="L257"></a>}

    <a id="L259"></a><span class="comment">// Otherwise, try suffixes.</span>
    <a id="L260"></a>for i := 0; i &lt; len(cfg.search); i++ {
        <a id="L261"></a>rname := name + &#34;.&#34; + cfg.search[i];
        <a id="L262"></a>if rname[len(rname)-1] != &#39;.&#39; {
            <a id="L263"></a>rname += &#34;.&#34;
        <a id="L264"></a>}
        <a id="L265"></a>addrs, err = tryOneName(cfg, rname);
        <a id="L266"></a>if err == nil {
            <a id="L267"></a>cname = rname;
            <a id="L268"></a>return;
        <a id="L269"></a>}
    <a id="L270"></a>}

    <a id="L272"></a><span class="comment">// Last ditch effort: try unsuffixed.</span>
    <a id="L273"></a>rname := name;
    <a id="L274"></a>if !rooted {
        <a id="L275"></a>rname += &#34;.&#34;
    <a id="L276"></a>}
    <a id="L277"></a>addrs, err = tryOneName(cfg, rname);
    <a id="L278"></a>if err == nil {
        <a id="L279"></a>cname = rname;
        <a id="L280"></a>return;
    <a id="L281"></a>}
    <a id="L282"></a>return;
<a id="L283"></a>}
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
