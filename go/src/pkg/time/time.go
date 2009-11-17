<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/time/time.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/time/time.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The time package provides functionality for measuring and</span>
<a id="L6"></a><span class="comment">// displaying time.</span>
<a id="L7"></a>package time

<a id="L9"></a>import (
    <a id="L10"></a>&#34;os&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">// Seconds reports the number of seconds since the Unix epoch,</span>
<a id="L14"></a><span class="comment">// January 1, 1970 00:00:00 UTC.</span>
<a id="L15"></a>func Seconds() int64 {
    <a id="L16"></a>sec, _, err := os.Time();
    <a id="L17"></a>if err != nil {
        <a id="L18"></a>panic(&#34;time: os.Time: &#34;, err.String())
    <a id="L19"></a>}
    <a id="L20"></a>return sec;
<a id="L21"></a>}

<a id="L23"></a><span class="comment">// Nanoseconds reports the number of nanoseconds since the Unix epoch,</span>
<a id="L24"></a><span class="comment">// January 1, 1970 00:00:00 UTC.</span>
<a id="L25"></a>func Nanoseconds() int64 {
    <a id="L26"></a>sec, nsec, err := os.Time();
    <a id="L27"></a>if err != nil {
        <a id="L28"></a>panic(&#34;time: os.Time: &#34;, err.String())
    <a id="L29"></a>}
    <a id="L30"></a>return sec*1e9 + nsec;
<a id="L31"></a>}

<a id="L33"></a><span class="comment">// Days of the week.</span>
<a id="L34"></a>const (
    <a id="L35"></a>Sunday = iota;
    <a id="L36"></a>Monday;
    <a id="L37"></a>Tuesday;
    <a id="L38"></a>Wednesday;
    <a id="L39"></a>Thursday;
    <a id="L40"></a>Friday;
    <a id="L41"></a>Saturday;
<a id="L42"></a>)

<a id="L44"></a><span class="comment">// Time is the struct representing a parsed time value.</span>
<a id="L45"></a>type Time struct {
    <a id="L46"></a>Year                 int64; <span class="comment">// 2008 is 2008</span>
    <a id="L47"></a>Month, Day           int;   <span class="comment">// Sep-17 is 9, 17</span>
    <a id="L48"></a>Hour, Minute, Second int;   <span class="comment">// 10:43:12 is 10, 43, 12</span>
    <a id="L49"></a>Weekday              int;   <span class="comment">// Sunday, Monday, ...</span>
    <a id="L50"></a>ZoneOffset           int;   <span class="comment">// seconds east of UTC</span>
    <a id="L51"></a>Zone                 string;
<a id="L52"></a>}

<a id="L54"></a>var nonleapyear = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
<a id="L55"></a>var leapyear = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

<a id="L57"></a>func months(year int64) []int {
    <a id="L58"></a>if year%4 == 0 &amp;&amp; (year%100 != 0 || year%400 == 0) {
        <a id="L59"></a>return leapyear
    <a id="L60"></a>}
    <a id="L61"></a>return nonleapyear;
<a id="L62"></a>}

<a id="L64"></a>const (
    <a id="L65"></a>secondsPerDay   = 24 * 60 * 60;
    <a id="L66"></a>daysPer400Years = 365*400 + 97;
    <a id="L67"></a>daysPer100Years = 365*100 + 24;
    <a id="L68"></a>daysPer4Years   = 365*4 + 1;
    <a id="L69"></a>days1970To2001  = 31*365 + 8;
<a id="L70"></a>)

<a id="L72"></a><span class="comment">// SecondsToUTC converts sec, in number of seconds since the Unix epoch,</span>
<a id="L73"></a><span class="comment">// into a parsed Time value in the UTC time zone.</span>
<a id="L74"></a>func SecondsToUTC(sec int64) *Time {
    <a id="L75"></a>t := new(Time);

    <a id="L77"></a><span class="comment">// Split into time and day.</span>
    <a id="L78"></a>day := sec / secondsPerDay;
    <a id="L79"></a>sec -= day * secondsPerDay;
    <a id="L80"></a>if sec &lt; 0 {
        <a id="L81"></a>day--;
        <a id="L82"></a>sec += secondsPerDay;
    <a id="L83"></a>}

    <a id="L85"></a><span class="comment">// Time</span>
    <a id="L86"></a>t.Hour = int(sec / 3600);
    <a id="L87"></a>t.Minute = int((sec / 60) % 60);
    <a id="L88"></a>t.Second = int(sec % 60);

    <a id="L90"></a><span class="comment">// Day 0 = January 1, 1970 was a Thursday</span>
    <a id="L91"></a>t.Weekday = int((day + Thursday) % 7);
    <a id="L92"></a>if t.Weekday &lt; 0 {
        <a id="L93"></a>t.Weekday += 7
    <a id="L94"></a>}

    <a id="L96"></a><span class="comment">// Change day from 0 = 1970 to 0 = 2001,</span>
    <a id="L97"></a><span class="comment">// to make leap year calculations easier</span>
    <a id="L98"></a><span class="comment">// (2001 begins 4-, 100-, and 400-year cycles ending in a leap year.)</span>
    <a id="L99"></a>day -= days1970To2001;

    <a id="L101"></a>year := int64(2001);
    <a id="L102"></a>if day &lt; 0 {
        <a id="L103"></a><span class="comment">// Go back enough 400 year cycles to make day positive.</span>
        <a id="L104"></a>n := -day/daysPer400Years + 1;
        <a id="L105"></a>year -= 400 * n;
        <a id="L106"></a>day += daysPer400Years * n;
    <a id="L107"></a>} else {
        <a id="L108"></a><span class="comment">// Cut off 400 year cycles.</span>
        <a id="L109"></a>n := day / daysPer400Years;
        <a id="L110"></a>year += 400 * n;
        <a id="L111"></a>day -= daysPer400Years * n;
    <a id="L112"></a>}

    <a id="L114"></a><span class="comment">// Cut off 100-year cycles</span>
    <a id="L115"></a>n := day / daysPer100Years;
    <a id="L116"></a>year += 100 * n;
    <a id="L117"></a>day -= daysPer100Years * n;

    <a id="L119"></a><span class="comment">// Cut off 4-year cycles</span>
    <a id="L120"></a>n = day / daysPer4Years;
    <a id="L121"></a>year += 4 * n;
    <a id="L122"></a>day -= daysPer4Years * n;

    <a id="L124"></a><span class="comment">// Cut off non-leap years.</span>
    <a id="L125"></a>n = day / 365;
    <a id="L126"></a>year += n;
    <a id="L127"></a>day -= 365 * n;

    <a id="L129"></a>t.Year = year;

    <a id="L131"></a><span class="comment">// If someone ever needs yearday,</span>
    <a id="L132"></a><span class="comment">// tyearday = day (+1?)</span>

    <a id="L134"></a>months := months(year);
    <a id="L135"></a>var m int;
    <a id="L136"></a>yday := int(day);
    <a id="L137"></a>for m = 0; m &lt; 12 &amp;&amp; yday &gt;= months[m]; m++ {
        <a id="L138"></a>yday -= months[m]
    <a id="L139"></a>}
    <a id="L140"></a>t.Month = m + 1;
    <a id="L141"></a>t.Day = yday + 1;
    <a id="L142"></a>t.Zone = &#34;UTC&#34;;

    <a id="L144"></a>return t;
<a id="L145"></a>}

<a id="L147"></a><span class="comment">// UTC returns the current time as a parsed Time value in the UTC time zone.</span>
<a id="L148"></a>func UTC() *Time { return SecondsToUTC(Seconds()) }

<a id="L150"></a><span class="comment">// SecondsToLocalTime converts sec, in number of seconds since the Unix epoch,</span>
<a id="L151"></a><span class="comment">// into a parsed Time value in the local time zone.</span>
<a id="L152"></a>func SecondsToLocalTime(sec int64) *Time {
    <a id="L153"></a>z, offset := lookupTimezone(sec);
    <a id="L154"></a>t := SecondsToUTC(sec + int64(offset));
    <a id="L155"></a>t.Zone = z;
    <a id="L156"></a>t.ZoneOffset = offset;
    <a id="L157"></a>return t;
<a id="L158"></a>}

<a id="L160"></a><span class="comment">// LocalTime returns the current time as a parsed Time value in the local time zone.</span>
<a id="L161"></a>func LocalTime() *Time { return SecondsToLocalTime(Seconds()) }

<a id="L163"></a><span class="comment">// Seconds returns the number of seconds since January 1, 1970 represented by the</span>
<a id="L164"></a><span class="comment">// parsed Time value.</span>
<a id="L165"></a>func (t *Time) Seconds() int64 {
    <a id="L166"></a><span class="comment">// First, accumulate days since January 1, 2001.</span>
    <a id="L167"></a><span class="comment">// Using 2001 instead of 1970 makes the leap-year</span>
    <a id="L168"></a><span class="comment">// handling easier (see SecondsToUTC), because</span>
    <a id="L169"></a><span class="comment">// it is at the beginning of the 4-, 100-, and 400-year cycles.</span>
    <a id="L170"></a>day := int64(0);

    <a id="L172"></a><span class="comment">// Rewrite year to be &gt;= 2001.</span>
    <a id="L173"></a>year := t.Year;
    <a id="L174"></a>if year &lt; 2001 {
        <a id="L175"></a>n := (2001-year)/400 + 1;
        <a id="L176"></a>year += 400 * n;
        <a id="L177"></a>day -= daysPer400Years * n;
    <a id="L178"></a>}

    <a id="L180"></a><span class="comment">// Add in days from 400-year cycles.</span>
    <a id="L181"></a>n := (year - 2001) / 400;
    <a id="L182"></a>year -= 400 * n;
    <a id="L183"></a>day += daysPer400Years * n;

    <a id="L185"></a><span class="comment">// Add in 100-year cycles.</span>
    <a id="L186"></a>n = (year - 2001) / 100;
    <a id="L187"></a>year -= 100 * n;
    <a id="L188"></a>day += daysPer100Years * n;

    <a id="L190"></a><span class="comment">// Add in 4-year cycles.</span>
    <a id="L191"></a>n = (year - 2001) / 4;
    <a id="L192"></a>year -= 4 * n;
    <a id="L193"></a>day += daysPer4Years * n;

    <a id="L195"></a><span class="comment">// Add in non-leap years.</span>
    <a id="L196"></a>n = year - 2001;
    <a id="L197"></a>day += 365 * n;

    <a id="L199"></a><span class="comment">// Add in days this year.</span>
    <a id="L200"></a>months := months(t.Year);
    <a id="L201"></a>for m := 0; m &lt; t.Month-1; m++ {
        <a id="L202"></a>day += int64(months[m])
    <a id="L203"></a>}
    <a id="L204"></a>day += int64(t.Day - 1);

    <a id="L206"></a><span class="comment">// Convert days to seconds since January 1, 2001.</span>
    <a id="L207"></a>sec := day * secondsPerDay;

    <a id="L209"></a><span class="comment">// Add in time elapsed today.</span>
    <a id="L210"></a>sec += int64(t.Hour) * 3600;
    <a id="L211"></a>sec += int64(t.Minute) * 60;
    <a id="L212"></a>sec += int64(t.Second);

    <a id="L214"></a><span class="comment">// Convert from seconds since 2001 to seconds since 1970.</span>
    <a id="L215"></a>sec += days1970To2001 * secondsPerDay;

    <a id="L217"></a><span class="comment">// Account for local time zone.</span>
    <a id="L218"></a>sec -= int64(t.ZoneOffset);
    <a id="L219"></a>return sec;
<a id="L220"></a>}

<a id="L222"></a>var longDayNames = []string{
    <a id="L223"></a>&#34;Sunday&#34;,
    <a id="L224"></a>&#34;Monday&#34;,
    <a id="L225"></a>&#34;Tuesday&#34;,
    <a id="L226"></a>&#34;Wednesday&#34;,
    <a id="L227"></a>&#34;Thursday&#34;,
    <a id="L228"></a>&#34;Friday&#34;,
    <a id="L229"></a>&#34;Saturday&#34;,
<a id="L230"></a>}

<a id="L232"></a>var shortDayNames = []string{
    <a id="L233"></a>&#34;Sun&#34;,
    <a id="L234"></a>&#34;Mon&#34;,
    <a id="L235"></a>&#34;Tue&#34;,
    <a id="L236"></a>&#34;Wed&#34;,
    <a id="L237"></a>&#34;Thu&#34;,
    <a id="L238"></a>&#34;Fri&#34;,
    <a id="L239"></a>&#34;Sat&#34;,
<a id="L240"></a>}

<a id="L242"></a>var shortMonthNames = []string{
    <a id="L243"></a>&#34;---&#34;,
    <a id="L244"></a>&#34;Jan&#34;,
    <a id="L245"></a>&#34;Feb&#34;,
    <a id="L246"></a>&#34;Mar&#34;,
    <a id="L247"></a>&#34;Apr&#34;,
    <a id="L248"></a>&#34;May&#34;,
    <a id="L249"></a>&#34;Jun&#34;,
    <a id="L250"></a>&#34;Jul&#34;,
    <a id="L251"></a>&#34;Aug&#34;,
    <a id="L252"></a>&#34;Sep&#34;,
    <a id="L253"></a>&#34;Oct&#34;,
    <a id="L254"></a>&#34;Nov&#34;,
    <a id="L255"></a>&#34;Dec&#34;,
<a id="L256"></a>}

<a id="L258"></a>func copy(dst []byte, s string) {
    <a id="L259"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L260"></a>dst[i] = s[i]
    <a id="L261"></a>}
<a id="L262"></a>}

<a id="L264"></a>func decimal(dst []byte, n int) {
    <a id="L265"></a>if n &lt; 0 {
        <a id="L266"></a>n = 0
    <a id="L267"></a>}
    <a id="L268"></a>for i := len(dst) - 1; i &gt;= 0; i-- {
        <a id="L269"></a>dst[i] = byte(n%10 + &#39;0&#39;);
        <a id="L270"></a>n /= 10;
    <a id="L271"></a>}
<a id="L272"></a>}

<a id="L274"></a>func addString(buf []byte, bp int, s string) int {
    <a id="L275"></a>n := len(s);
    <a id="L276"></a>copy(buf[bp:bp+n], s);
    <a id="L277"></a>return bp + n;
<a id="L278"></a>}

<a id="L280"></a><span class="comment">// Just enough of strftime to implement the date formats below.</span>
<a id="L281"></a><span class="comment">// Not exported.</span>
<a id="L282"></a>func format(t *Time, fmt string) string {
    <a id="L283"></a>buf := make([]byte, 128);
    <a id="L284"></a>bp := 0;

    <a id="L286"></a>for i := 0; i &lt; len(fmt); i++ {
        <a id="L287"></a>if fmt[i] == &#39;%&#39; {
            <a id="L288"></a>i++;
            <a id="L289"></a>switch fmt[i] {
            <a id="L290"></a>case &#39;A&#39;: <span class="comment">// %A full weekday name</span>
                <a id="L291"></a>bp = addString(buf, bp, longDayNames[t.Weekday])
            <a id="L292"></a>case &#39;a&#39;: <span class="comment">// %a abbreviated weekday name</span>
                <a id="L293"></a>bp = addString(buf, bp, shortDayNames[t.Weekday])
            <a id="L294"></a>case &#39;b&#39;: <span class="comment">// %b abbreviated month name</span>
                <a id="L295"></a>bp = addString(buf, bp, shortMonthNames[t.Month])
            <a id="L296"></a>case &#39;d&#39;: <span class="comment">// %d day of month (01-31)</span>
                <a id="L297"></a>decimal(buf[bp:bp+2], t.Day);
                <a id="L298"></a>bp += 2;
            <a id="L299"></a>case &#39;e&#39;: <span class="comment">// %e day of month ( 1-31)</span>
                <a id="L300"></a>if t.Day &gt;= 10 {
                    <a id="L301"></a>decimal(buf[bp:bp+2], t.Day)
                <a id="L302"></a>} else {
                    <a id="L303"></a>buf[bp] = &#39; &#39;;
                    <a id="L304"></a>buf[bp+1] = byte(t.Day + &#39;0&#39;);
                <a id="L305"></a>}
                <a id="L306"></a>bp += 2;
            <a id="L307"></a>case &#39;H&#39;: <span class="comment">// %H hour 00-23</span>
                <a id="L308"></a>decimal(buf[bp:bp+2], t.Hour);
                <a id="L309"></a>bp += 2;
            <a id="L310"></a>case &#39;M&#39;: <span class="comment">// %M minute 00-59</span>
                <a id="L311"></a>decimal(buf[bp:bp+2], t.Minute);
                <a id="L312"></a>bp += 2;
            <a id="L313"></a>case &#39;S&#39;: <span class="comment">// %S second 00-59</span>
                <a id="L314"></a>decimal(buf[bp:bp+2], t.Second);
                <a id="L315"></a>bp += 2;
            <a id="L316"></a>case &#39;Y&#39;: <span class="comment">// %Y year 2008</span>
                <a id="L317"></a>decimal(buf[bp:bp+4], int(t.Year));
                <a id="L318"></a>bp += 4;
            <a id="L319"></a>case &#39;y&#39;: <span class="comment">// %y year 08</span>
                <a id="L320"></a>decimal(buf[bp:bp+2], int(t.Year%100));
                <a id="L321"></a>bp += 2;
            <a id="L322"></a>case &#39;Z&#39;:
                <a id="L323"></a>bp = addString(buf, bp, t.Zone)
            <a id="L324"></a>default:
                <a id="L325"></a>buf[bp] = &#39;%&#39;;
                <a id="L326"></a>buf[bp+1] = fmt[i];
                <a id="L327"></a>bp += 2;
            <a id="L328"></a>}
        <a id="L329"></a>} else {
            <a id="L330"></a>buf[bp] = fmt[i];
            <a id="L331"></a>bp++;
        <a id="L332"></a>}
    <a id="L333"></a>}
    <a id="L334"></a>return string(buf[0:bp]);
<a id="L335"></a>}

<a id="L337"></a><span class="comment">// Asctime formats the parsed time value in the style of</span>
<a id="L338"></a><span class="comment">// ANSI C asctime: Sun Nov  6 08:49:37 1994</span>
<a id="L339"></a>func (t *Time) Asctime() string { return format(t, &#34;%a %b %e %H:%M:%S %Y&#34;) }

<a id="L341"></a><span class="comment">// RFC850 formats the parsed time value in the style of</span>
<a id="L342"></a><span class="comment">// RFC 850: Sunday, 06-Nov-94 08:49:37 UTC</span>
<a id="L343"></a>func (t *Time) RFC850() string { return format(t, &#34;%A, %d-%b-%y %H:%M:%S %Z&#34;) }

<a id="L345"></a><span class="comment">// RFC1123 formats the parsed time value in the style of</span>
<a id="L346"></a><span class="comment">// RFC 1123: Sun, 06 Nov 1994 08:49:37 UTC</span>
<a id="L347"></a>func (t *Time) RFC1123() string { return format(t, &#34;%a, %d %b %Y %H:%M:%S %Z&#34;) }

<a id="L349"></a><span class="comment">// String formats the parsed time value in the style of</span>
<a id="L350"></a><span class="comment">// date(1) - Sun Nov  6 08:49:37 UTC 1994</span>
<a id="L351"></a>func (t *Time) String() string { return format(t, &#34;%a %b %e %H:%M:%S %Z %Y&#34;) }
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
