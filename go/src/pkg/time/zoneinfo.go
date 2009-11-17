<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/time/zoneinfo.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/time/zoneinfo.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Parse &#34;zoneinfo&#34; time zone file.</span>
<a id="L6"></a><span class="comment">// This is a fairly standard file format used on OS X, Linux, BSD, Sun, and others.</span>
<a id="L7"></a><span class="comment">// See tzfile(5), http://en.wikipedia.org/wiki/Zoneinfo,</span>
<a id="L8"></a><span class="comment">// and ftp://munnari.oz.au/pub/oldtz/</span>

<a id="L10"></a>package time

<a id="L12"></a>import (
    <a id="L13"></a>&#34;io&#34;;
    <a id="L14"></a>&#34;once&#34;;
    <a id="L15"></a>&#34;os&#34;;
<a id="L16"></a>)

<a id="L18"></a>const (
    <a id="L19"></a>headerSize = 4 + 16 + 4*7;
    <a id="L20"></a>zoneDir    = &#34;/usr/share/zoneinfo/&#34;;
<a id="L21"></a>)

<a id="L23"></a><span class="comment">// Simple I/O interface to binary blob of data.</span>
<a id="L24"></a>type data struct {
    <a id="L25"></a>p     []byte;
    <a id="L26"></a>error bool;
<a id="L27"></a>}


<a id="L30"></a>func (d *data) read(n int) []byte {
    <a id="L31"></a>if len(d.p) &lt; n {
        <a id="L32"></a>d.p = nil;
        <a id="L33"></a>d.error = true;
        <a id="L34"></a>return nil;
    <a id="L35"></a>}
    <a id="L36"></a>p := d.p[0:n];
    <a id="L37"></a>d.p = d.p[n:len(d.p)];
    <a id="L38"></a>return p;
<a id="L39"></a>}

<a id="L41"></a>func (d *data) big4() (n uint32, ok bool) {
    <a id="L42"></a>p := d.read(4);
    <a id="L43"></a>if len(p) &lt; 4 {
        <a id="L44"></a>d.error = true;
        <a id="L45"></a>return 0, false;
    <a id="L46"></a>}
    <a id="L47"></a>return uint32(p[0])&lt;&lt;24 | uint32(p[1])&lt;&lt;16 | uint32(p[2])&lt;&lt;8 | uint32(p[3]), true;
<a id="L48"></a>}

<a id="L50"></a>func (d *data) byte() (n byte, ok bool) {
    <a id="L51"></a>p := d.read(1);
    <a id="L52"></a>if len(p) &lt; 1 {
        <a id="L53"></a>d.error = true;
        <a id="L54"></a>return 0, false;
    <a id="L55"></a>}
    <a id="L56"></a>return p[0], true;
<a id="L57"></a>}


<a id="L60"></a><span class="comment">// Make a string by stopping at the first NUL</span>
<a id="L61"></a>func byteString(p []byte) string {
    <a id="L62"></a>for i := 0; i &lt; len(p); i++ {
        <a id="L63"></a>if p[i] == 0 {
            <a id="L64"></a>return string(p[0:i])
        <a id="L65"></a>}
    <a id="L66"></a>}
    <a id="L67"></a>return string(p);
<a id="L68"></a>}

<a id="L70"></a><span class="comment">// Parsed representation</span>
<a id="L71"></a>type zone struct {
    <a id="L72"></a>utcoff int;
    <a id="L73"></a>isdst  bool;
    <a id="L74"></a>name   string;
<a id="L75"></a>}

<a id="L77"></a>type zonetime struct {
    <a id="L78"></a>time         int32; <span class="comment">// transition time, in seconds since 1970 GMT</span>
    <a id="L79"></a>zone         *zone; <span class="comment">// the zone that goes into effect at that time</span>
    <a id="L80"></a>isstd, isutc bool;  <span class="comment">// ignored - no idea what these mean</span>
<a id="L81"></a>}

<a id="L83"></a>func parseinfo(bytes []byte) (zt []zonetime, ok bool) {
    <a id="L84"></a>d := data{bytes, false};

    <a id="L86"></a><span class="comment">// 4-byte magic &#34;TZif&#34;</span>
    <a id="L87"></a>if magic := d.read(4); string(magic) != &#34;TZif&#34; {
        <a id="L88"></a>return nil, false
    <a id="L89"></a>}

    <a id="L91"></a><span class="comment">// 1-byte version, then 15 bytes of padding</span>
    <a id="L92"></a>var p []byte;
    <a id="L93"></a>if p = d.read(16); len(p) != 16 || p[0] != 0 &amp;&amp; p[0] != &#39;2&#39; {
        <a id="L94"></a>return nil, false
    <a id="L95"></a>}

    <a id="L97"></a><span class="comment">// six big-endian 32-bit integers:</span>
    <a id="L98"></a><span class="comment">//	number of UTC/local indicators</span>
    <a id="L99"></a><span class="comment">//	number of standard/wall indicators</span>
    <a id="L100"></a><span class="comment">//	number of leap seconds</span>
    <a id="L101"></a><span class="comment">//	number of transition times</span>
    <a id="L102"></a><span class="comment">//	number of local time zones</span>
    <a id="L103"></a><span class="comment">//	number of characters of time zone abbrev strings</span>
    <a id="L104"></a>const (
        <a id="L105"></a>NUTCLocal = iota;
        <a id="L106"></a>NStdWall;
        <a id="L107"></a>NLeap;
        <a id="L108"></a>NTime;
        <a id="L109"></a>NZone;
        <a id="L110"></a>NChar;
    <a id="L111"></a>)
    <a id="L112"></a>var n [6]int;
    <a id="L113"></a>for i := 0; i &lt; 6; i++ {
        <a id="L114"></a>nn, ok := d.big4();
        <a id="L115"></a>if !ok {
            <a id="L116"></a>return nil, false
        <a id="L117"></a>}
        <a id="L118"></a>n[i] = int(nn);
    <a id="L119"></a>}

    <a id="L121"></a><span class="comment">// Transition times.</span>
    <a id="L122"></a>txtimes := data{d.read(n[NTime] * 4), false};

    <a id="L124"></a><span class="comment">// Time zone indices for transition times.</span>
    <a id="L125"></a>txzones := d.read(n[NTime]);

    <a id="L127"></a><span class="comment">// Zone info structures</span>
    <a id="L128"></a>zonedata := data{d.read(n[NZone] * 6), false};

    <a id="L130"></a><span class="comment">// Time zone abbreviations.</span>
    <a id="L131"></a>abbrev := d.read(n[NChar]);

    <a id="L133"></a><span class="comment">// Leap-second time pairs</span>
    <a id="L134"></a>d.read(n[NLeap] * 8);

    <a id="L136"></a><span class="comment">// Whether tx times associated with local time types</span>
    <a id="L137"></a><span class="comment">// are specified as standard time or wall time.</span>
    <a id="L138"></a>isstd := d.read(n[NStdWall]);

    <a id="L140"></a><span class="comment">// Whether tx times associated with local time types</span>
    <a id="L141"></a><span class="comment">// are specified as UTC or local time.</span>
    <a id="L142"></a>isutc := d.read(n[NUTCLocal]);

    <a id="L144"></a>if d.error { <span class="comment">// ran out of data</span>
        <a id="L145"></a>return nil, false
    <a id="L146"></a>}

    <a id="L148"></a><span class="comment">// If version == 2, the entire file repeats, this time using</span>
    <a id="L149"></a><span class="comment">// 8-byte ints for txtimes and leap seconds.</span>
    <a id="L150"></a><span class="comment">// We won&#39;t need those until 2106.</span>

    <a id="L152"></a><span class="comment">// Now we can build up a useful data structure.</span>
    <a id="L153"></a><span class="comment">// First the zone information.</span>
    <a id="L154"></a><span class="comment">//	utcoff[4] isdst[1] nameindex[1]</span>
    <a id="L155"></a>z := make([]zone, n[NZone]);
    <a id="L156"></a>for i := 0; i &lt; len(z); i++ {
        <a id="L157"></a>var ok bool;
        <a id="L158"></a>var n uint32;
        <a id="L159"></a>if n, ok = zonedata.big4(); !ok {
            <a id="L160"></a>return nil, false
        <a id="L161"></a>}
        <a id="L162"></a>z[i].utcoff = int(n);
        <a id="L163"></a>var b byte;
        <a id="L164"></a>if b, ok = zonedata.byte(); !ok {
            <a id="L165"></a>return nil, false
        <a id="L166"></a>}
        <a id="L167"></a>z[i].isdst = b != 0;
        <a id="L168"></a>if b, ok = zonedata.byte(); !ok || int(b) &gt;= len(abbrev) {
            <a id="L169"></a>return nil, false
        <a id="L170"></a>}
        <a id="L171"></a>z[i].name = byteString(abbrev[b:len(abbrev)]);
    <a id="L172"></a>}

    <a id="L174"></a><span class="comment">// Now the transition time info.</span>
    <a id="L175"></a>zt = make([]zonetime, n[NTime]);
    <a id="L176"></a>for i := 0; i &lt; len(zt); i++ {
        <a id="L177"></a>var ok bool;
        <a id="L178"></a>var n uint32;
        <a id="L179"></a>if n, ok = txtimes.big4(); !ok {
            <a id="L180"></a>return nil, false
        <a id="L181"></a>}
        <a id="L182"></a>zt[i].time = int32(n);
        <a id="L183"></a>if int(txzones[i]) &gt;= len(z) {
            <a id="L184"></a>return nil, false
        <a id="L185"></a>}
        <a id="L186"></a>zt[i].zone = &amp;z[txzones[i]];
        <a id="L187"></a>if i &lt; len(isstd) {
            <a id="L188"></a>zt[i].isstd = isstd[i] != 0
        <a id="L189"></a>}
        <a id="L190"></a>if i &lt; len(isutc) {
            <a id="L191"></a>zt[i].isutc = isutc[i] != 0
        <a id="L192"></a>}
    <a id="L193"></a>}
    <a id="L194"></a>return zt, true;
<a id="L195"></a>}

<a id="L197"></a>func readinfofile(name string) ([]zonetime, bool) {
    <a id="L198"></a>buf, err := io.ReadFile(name);
    <a id="L199"></a>if err != nil {
        <a id="L200"></a>return nil, false
    <a id="L201"></a>}
    <a id="L202"></a>return parseinfo(buf);
<a id="L203"></a>}

<a id="L205"></a>var zones []zonetime

<a id="L207"></a>func setupZone() {
    <a id="L208"></a><span class="comment">// consult $TZ to find the time zone to use.</span>
    <a id="L209"></a><span class="comment">// no $TZ means use the system default /etc/localtime.</span>
    <a id="L210"></a><span class="comment">// $TZ=&#34;&#34; means use UTC.</span>
    <a id="L211"></a><span class="comment">// $TZ=&#34;foo&#34; means use /usr/share/zoneinfo/foo.</span>

    <a id="L213"></a>tz, err := os.Getenverror(&#34;TZ&#34;);
    <a id="L214"></a>switch {
    <a id="L215"></a>case err == os.ENOENV:
        <a id="L216"></a>zones, _ = readinfofile(&#34;/etc/localtime&#34;)
    <a id="L217"></a>case len(tz) &gt; 0:
        <a id="L218"></a>zones, _ = readinfofile(zoneDir + tz)
    <a id="L219"></a>case len(tz) == 0:
        <a id="L220"></a><span class="comment">// do nothing: use UTC</span>
    <a id="L221"></a>}
<a id="L222"></a>}

<a id="L224"></a>func lookupTimezone(sec int64) (zone string, offset int) {
    <a id="L225"></a>once.Do(setupZone);
    <a id="L226"></a>if len(zones) == 0 {
        <a id="L227"></a>return &#34;UTC&#34;, 0
    <a id="L228"></a>}

    <a id="L230"></a><span class="comment">// Binary search for entry with largest time &lt;= sec</span>
    <a id="L231"></a>tz := zones;
    <a id="L232"></a>for len(tz) &gt; 1 {
        <a id="L233"></a>m := len(tz) / 2;
        <a id="L234"></a>if sec &lt; int64(tz[m].time) {
            <a id="L235"></a>tz = tz[0:m]
        <a id="L236"></a>} else {
            <a id="L237"></a>tz = tz[m:len(tz)]
        <a id="L238"></a>}
    <a id="L239"></a>}
    <a id="L240"></a>z := tz[0].zone;
    <a id="L241"></a>return z.name, z.utcoff;
<a id="L242"></a>}
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
