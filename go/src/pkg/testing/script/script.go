<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/testing/script/script.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/testing/script/script.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package aids in the testing of code that uses channels.</span>
<a id="L6"></a>package script

<a id="L8"></a>import (
    <a id="L9"></a>&#34;fmt&#34;;
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;rand&#34;;
    <a id="L12"></a>&#34;reflect&#34;;
    <a id="L13"></a>&#34;strings&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// An Event is an element in a partially ordered set that either sends a value</span>
<a id="L17"></a><span class="comment">// to a channel or expects a value from a channel.</span>
<a id="L18"></a>type Event struct {
    <a id="L19"></a>name         string;
    <a id="L20"></a>occurred     bool;
    <a id="L21"></a>predecessors []*Event;
    <a id="L22"></a>action       action;
<a id="L23"></a>}

<a id="L25"></a>type action interface {
    <a id="L26"></a><span class="comment">// getSend returns nil if the action is not a send action.</span>
    <a id="L27"></a>getSend() sendAction;
    <a id="L28"></a><span class="comment">// getRecv returns nil if the action is not a receive action.</span>
    <a id="L29"></a>getRecv() recvAction;
    <a id="L30"></a><span class="comment">// getChannel returns the channel that the action operates on.</span>
    <a id="L31"></a>getChannel() interface{};
<a id="L32"></a>}

<a id="L34"></a>type recvAction interface {
    <a id="L35"></a>recvMatch(interface{}) bool;
<a id="L36"></a>}

<a id="L38"></a>type sendAction interface {
    <a id="L39"></a>send();
<a id="L40"></a>}

<a id="L42"></a><span class="comment">// isReady returns true if all the predecessors of an Event have occurred.</span>
<a id="L43"></a>func (e Event) isReady() bool {
    <a id="L44"></a>for _, predecessor := range e.predecessors {
        <a id="L45"></a>if !predecessor.occurred {
            <a id="L46"></a>return false
        <a id="L47"></a>}
    <a id="L48"></a>}

    <a id="L50"></a>return true;
<a id="L51"></a>}

<a id="L53"></a><span class="comment">// A Recv action reads a value from a channel and uses reflect.DeepMatch to</span>
<a id="L54"></a><span class="comment">// compare it with an expected value.</span>
<a id="L55"></a>type Recv struct {
    <a id="L56"></a>Channel  interface{};
    <a id="L57"></a>Expected interface{};
<a id="L58"></a>}

<a id="L60"></a>func (r Recv) getRecv() recvAction { return r }

<a id="L62"></a>func (Recv) getSend() sendAction { return nil }

<a id="L64"></a>func (r Recv) getChannel() interface{} { return r.Channel }

<a id="L66"></a>func (r Recv) recvMatch(chanEvent interface{}) bool {
    <a id="L67"></a>c, ok := chanEvent.(channelRecv);
    <a id="L68"></a>if !ok || c.channel != r.Channel {
        <a id="L69"></a>return false
    <a id="L70"></a>}

    <a id="L72"></a>return reflect.DeepEqual(c.value, r.Expected);
<a id="L73"></a>}

<a id="L75"></a><span class="comment">// A RecvMatch action reads a value from a channel and calls a function to</span>
<a id="L76"></a><span class="comment">// determine if the value matches.</span>
<a id="L77"></a>type RecvMatch struct {
    <a id="L78"></a>Channel interface{};
    <a id="L79"></a>Match   func(interface{}) bool;
<a id="L80"></a>}

<a id="L82"></a>func (r RecvMatch) getRecv() recvAction { return r }

<a id="L84"></a>func (RecvMatch) getSend() sendAction { return nil }

<a id="L86"></a>func (r RecvMatch) getChannel() interface{} { return r.Channel }

<a id="L88"></a>func (r RecvMatch) recvMatch(chanEvent interface{}) bool {
    <a id="L89"></a>c, ok := chanEvent.(channelRecv);
    <a id="L90"></a>if !ok || c.channel != r.Channel {
        <a id="L91"></a>return false
    <a id="L92"></a>}

    <a id="L94"></a>return r.Match(c.value);
<a id="L95"></a>}

<a id="L97"></a><span class="comment">// A Closed action matches if the given channel is closed. The closing is</span>
<a id="L98"></a><span class="comment">// treated as an event, not a state, thus Closed will only match once for a</span>
<a id="L99"></a><span class="comment">// given channel.</span>
<a id="L100"></a>type Closed struct {
    <a id="L101"></a>Channel interface{};
<a id="L102"></a>}

<a id="L104"></a>func (r Closed) getRecv() recvAction { return r }

<a id="L106"></a>func (Closed) getSend() sendAction { return nil }

<a id="L108"></a>func (r Closed) getChannel() interface{} { return r.Channel }

<a id="L110"></a>func (r Closed) recvMatch(chanEvent interface{}) bool {
    <a id="L111"></a>c, ok := chanEvent.(channelClosed);
    <a id="L112"></a>if !ok || c.channel != r.Channel {
        <a id="L113"></a>return false
    <a id="L114"></a>}

    <a id="L116"></a>return true;
<a id="L117"></a>}

<a id="L119"></a><span class="comment">// A Send action sends a value to a channel. The value must match the</span>
<a id="L120"></a><span class="comment">// type of the channel exactly unless the channel if of type chan interface{}.</span>
<a id="L121"></a>type Send struct {
    <a id="L122"></a>Channel interface{};
    <a id="L123"></a>Value   interface{};
<a id="L124"></a>}

<a id="L126"></a>func (Send) getRecv() recvAction { return nil }

<a id="L128"></a>func (s Send) getSend() sendAction { return s }

<a id="L130"></a>func (s Send) getChannel() interface{} { return s.Channel }

<a id="L132"></a>func newEmptyInterface(args ...) reflect.Value {
    <a id="L133"></a>return reflect.NewValue(args).(*reflect.StructValue).Field(0)
<a id="L134"></a>}

<a id="L136"></a>func (s Send) send() {
    <a id="L137"></a><span class="comment">// With reflect.ChanValue.Send, we must match the types exactly. So, if</span>
    <a id="L138"></a><span class="comment">// s.Channel is a chan interface{} we convert s.Value to an interface{}</span>
    <a id="L139"></a><span class="comment">// first.</span>
    <a id="L140"></a>c := reflect.NewValue(s.Channel).(*reflect.ChanValue);
    <a id="L141"></a>var v reflect.Value;
    <a id="L142"></a>if iface, ok := c.Type().(*reflect.ChanType).Elem().(*reflect.InterfaceType); ok &amp;&amp; iface.NumMethod() == 0 {
        <a id="L143"></a>v = newEmptyInterface(s.Value)
    <a id="L144"></a>} else {
        <a id="L145"></a>v = reflect.NewValue(s.Value)
    <a id="L146"></a>}
    <a id="L147"></a>c.Send(v);
<a id="L148"></a>}

<a id="L150"></a><span class="comment">// A Close action closes the given channel.</span>
<a id="L151"></a>type Close struct {
    <a id="L152"></a>Channel interface{};
<a id="L153"></a>}

<a id="L155"></a>func (Close) getRecv() recvAction { return nil }

<a id="L157"></a>func (s Close) getSend() sendAction { return s }

<a id="L159"></a>func (s Close) getChannel() interface{} { return s.Channel }

<a id="L161"></a>func (s Close) send() { reflect.NewValue(s.Channel).(*reflect.ChanValue).Close() }

<a id="L163"></a><span class="comment">// A ReceivedUnexpected error results if no active Events match a value</span>
<a id="L164"></a><span class="comment">// received from a channel.</span>
<a id="L165"></a>type ReceivedUnexpected struct {
    <a id="L166"></a>Value interface{};
    <a id="L167"></a>ready []*Event;
<a id="L168"></a>}

<a id="L170"></a>func (r ReceivedUnexpected) String() string {
    <a id="L171"></a>names := make([]string, len(r.ready));
    <a id="L172"></a>for i, v := range r.ready {
        <a id="L173"></a>names[i] = v.name
    <a id="L174"></a>}
    <a id="L175"></a>return fmt.Sprintf(&#34;received unexpected value on one of the channels: %#v. Runnable events: %s&#34;, r.Value, strings.Join(names, &#34;, &#34;));
<a id="L176"></a>}

<a id="L178"></a><span class="comment">// A SetupError results if there is a error with the configuration of a set of</span>
<a id="L179"></a><span class="comment">// Events.</span>
<a id="L180"></a>type SetupError string

<a id="L182"></a>func (s SetupError) String() string { return string(s) }

<a id="L184"></a>func NewEvent(name string, predecessors []*Event, action action) *Event {
    <a id="L185"></a>e := &amp;Event{name, false, predecessors, action};
    <a id="L186"></a>return e;
<a id="L187"></a>}

<a id="L189"></a><span class="comment">// Given a set of Events, Perform repeatedly iterates over the set and finds the</span>
<a id="L190"></a><span class="comment">// subset of ready Events (that is, all of their predecessors have</span>
<a id="L191"></a><span class="comment">// occurred). From that subset, it pseudo-randomly selects an Event to perform.</span>
<a id="L192"></a><span class="comment">// If the Event is a send event, the send occurs and Perform recalculates the ready</span>
<a id="L193"></a><span class="comment">// set. If the event is a receive event, Perform waits for a value from any of the</span>
<a id="L194"></a><span class="comment">// channels that are contained in any of the events. That value is then matched</span>
<a id="L195"></a><span class="comment">// against the ready events. The first event that matches is considered to</span>
<a id="L196"></a><span class="comment">// have occurred and Perform recalculates the ready set.</span>
<a id="L197"></a><span class="comment">//</span>
<a id="L198"></a><span class="comment">// Perform continues this until all Events have occurred.</span>
<a id="L199"></a><span class="comment">//</span>
<a id="L200"></a><span class="comment">// Note that uncollected goroutines may still be reading from any of the</span>
<a id="L201"></a><span class="comment">// channels read from after Perform returns.</span>
<a id="L202"></a><span class="comment">//</span>
<a id="L203"></a><span class="comment">// For example, consider the problem of testing a function that reads values on</span>
<a id="L204"></a><span class="comment">// one channel and echos them to two output channels. To test this we would</span>
<a id="L205"></a><span class="comment">// create three events: a send event and two receive events. Each of the</span>
<a id="L206"></a><span class="comment">// receive events must list the send event as a predecessor but there is no</span>
<a id="L207"></a><span class="comment">// ordering between the receive events.</span>
<a id="L208"></a><span class="comment">//</span>
<a id="L209"></a><span class="comment">//  send := NewEvent(&#34;send&#34;, nil, Send{c, 1});</span>
<a id="L210"></a><span class="comment">//  recv1 := NewEvent(&#34;recv 1&#34;, []*Event{send}, Recv{c, 1});</span>
<a id="L211"></a><span class="comment">//  recv2 := NewEvent(&#34;recv 2&#34;, []*Event{send}, Recv{c, 1});</span>
<a id="L212"></a><span class="comment">//  Perform(0, []*Event{send, recv1, recv2});</span>
<a id="L213"></a><span class="comment">//</span>
<a id="L214"></a><span class="comment">// At first, only the send event would be in the ready set and thus Perform will</span>
<a id="L215"></a><span class="comment">// send a value to the input channel. Now the two receive events are ready and</span>
<a id="L216"></a><span class="comment">// Perform will match each of them against the values read from the output channels.</span>
<a id="L217"></a><span class="comment">//</span>
<a id="L218"></a><span class="comment">// It would be invalid to list one of the receive events as a predecessor of</span>
<a id="L219"></a><span class="comment">// the other. At each receive step, all the receive channels are considered,</span>
<a id="L220"></a><span class="comment">// thus Perform may see a value from a channel that is not in the current ready</span>
<a id="L221"></a><span class="comment">// set and fail.</span>
<a id="L222"></a>func Perform(seed int64, events []*Event) (err os.Error) {
    <a id="L223"></a>r := rand.New(rand.NewSource(seed));

    <a id="L225"></a>channels, err := getChannels(events);
    <a id="L226"></a>if err != nil {
        <a id="L227"></a>return
    <a id="L228"></a>}
    <a id="L229"></a>multiplex := make(chan interface{});
    <a id="L230"></a>for _, channel := range channels {
        <a id="L231"></a>go recvValues(multiplex, channel)
    <a id="L232"></a>}

<a id="L234"></a>Outer:
    <a id="L235"></a>for {
        <a id="L236"></a>ready, err := readyEvents(events);
        <a id="L237"></a>if err != nil {
            <a id="L238"></a>return err
        <a id="L239"></a>}

        <a id="L241"></a>if len(ready) == 0 {
            <a id="L242"></a><span class="comment">// All events occurred.</span>
            <a id="L243"></a>break
        <a id="L244"></a>}

        <a id="L246"></a>event := ready[r.Intn(len(ready))];
        <a id="L247"></a>if send := event.action.getSend(); send != nil {
            <a id="L248"></a>send.send();
            <a id="L249"></a>event.occurred = true;
            <a id="L250"></a>continue;
        <a id="L251"></a>}

        <a id="L253"></a>v := &lt;-multiplex;
        <a id="L254"></a>for _, event := range ready {
            <a id="L255"></a>if recv := event.action.getRecv(); recv != nil &amp;&amp; recv.recvMatch(v) {
                <a id="L256"></a>event.occurred = true;
                <a id="L257"></a>continue Outer;
            <a id="L258"></a>}
        <a id="L259"></a>}

        <a id="L261"></a>return ReceivedUnexpected{v, ready};
    <a id="L262"></a>}

    <a id="L264"></a>return nil;
<a id="L265"></a>}

<a id="L267"></a><span class="comment">// getChannels returns all the channels listed in any receive events.</span>
<a id="L268"></a>func getChannels(events []*Event) ([]interface{}, os.Error) {
    <a id="L269"></a>channels := make([]interface{}, len(events));

    <a id="L271"></a>j := 0;
    <a id="L272"></a>for _, event := range events {
        <a id="L273"></a>if recv := event.action.getRecv(); recv == nil {
            <a id="L274"></a>continue
        <a id="L275"></a>}
        <a id="L276"></a>c := event.action.getChannel();
        <a id="L277"></a>if _, ok := reflect.NewValue(c).(*reflect.ChanValue); !ok {
            <a id="L278"></a>return nil, SetupError(&#34;one of the channel values is not a channel&#34;)
        <a id="L279"></a>}

        <a id="L281"></a>duplicate := false;
        <a id="L282"></a>for _, other := range channels[0:j] {
            <a id="L283"></a>if c == other {
                <a id="L284"></a>duplicate = true;
                <a id="L285"></a>break;
            <a id="L286"></a>}
        <a id="L287"></a>}

        <a id="L289"></a>if !duplicate {
            <a id="L290"></a>channels[j] = c;
            <a id="L291"></a>j++;
        <a id="L292"></a>}
    <a id="L293"></a>}

    <a id="L295"></a>return channels[0:j], nil;
<a id="L296"></a>}

<a id="L298"></a><span class="comment">// recvValues is a multiplexing helper function. It reads values from the given</span>
<a id="L299"></a><span class="comment">// channel repeatedly, wrapping them up as either a channelRecv or</span>
<a id="L300"></a><span class="comment">// channelClosed structure, and forwards them to the multiplex channel.</span>
<a id="L301"></a>func recvValues(multiplex chan&lt;- interface{}, channel interface{}) {
    <a id="L302"></a>c := reflect.NewValue(channel).(*reflect.ChanValue);

    <a id="L304"></a>for {
        <a id="L305"></a>v := c.Recv();
        <a id="L306"></a>if c.Closed() {
            <a id="L307"></a>multiplex &lt;- channelClosed{channel};
            <a id="L308"></a>return;
        <a id="L309"></a>}

        <a id="L311"></a>multiplex &lt;- channelRecv{channel, v.Interface()};
    <a id="L312"></a>}
<a id="L313"></a>}

<a id="L315"></a>type channelClosed struct {
    <a id="L316"></a>channel interface{};
<a id="L317"></a>}

<a id="L319"></a>type channelRecv struct {
    <a id="L320"></a>channel interface{};
    <a id="L321"></a>value   interface{};
<a id="L322"></a>}

<a id="L324"></a><span class="comment">// readyEvents returns the subset of events that are ready.</span>
<a id="L325"></a>func readyEvents(events []*Event) ([]*Event, os.Error) {
    <a id="L326"></a>ready := make([]*Event, len(events));

    <a id="L328"></a>j := 0;
    <a id="L329"></a>eventsWaiting := false;
    <a id="L330"></a>for _, event := range events {
        <a id="L331"></a>if event.occurred {
            <a id="L332"></a>continue
        <a id="L333"></a>}

        <a id="L335"></a>eventsWaiting = true;
        <a id="L336"></a>if event.isReady() {
            <a id="L337"></a>ready[j] = event;
            <a id="L338"></a>j++;
        <a id="L339"></a>}
    <a id="L340"></a>}

    <a id="L342"></a>if j == 0 &amp;&amp; eventsWaiting {
        <a id="L343"></a>names := make([]string, len(events));
        <a id="L344"></a>for _, event := range events {
            <a id="L345"></a>if event.occurred {
                <a id="L346"></a>continue
            <a id="L347"></a>}
            <a id="L348"></a>names[j] = event.name;
        <a id="L349"></a>}

        <a id="L351"></a>return nil, SetupError(&#34;dependency cycle in events. These events are waiting to run but cannot: &#34; + strings.Join(names, &#34;, &#34;));
    <a id="L352"></a>}

    <a id="L354"></a>return ready[0:j], nil;
<a id="L355"></a>}
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
