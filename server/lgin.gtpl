<html>
    <head>
    <title></title>
    </head>
    <body>
        <h1>Struct under construction: </h1>
        <form action="/" method="post">
            <ul>
              <li>WebsiteUrl: <input type="text" name="websiteurl" value={{.WebsiteUrl}}> </li>
              <li>SessionId:  <input type="text" name="sessionid" value={{.SessionId}}> </li>
              <li>ResizeFrom:<br/>  Width:<input type="text" name="resizefromwidth" value={{.ResizeFrom.Width}}>
                                    Height:<input type="text" name="resizefromheight" value={{.ResizeFrom.Height}}> </li>
              <li>ResizeTo:<br/>  Width:<input type="text" name="resizetowidth" value={{.ResizeTo.Width}}>
                                  Height:<input type="text" name="resizetoheight" value={{.ResizeTo.Height}}> </li>
              <li>CopyAndPaste:<br />
                  inputEmail: <select name="cpinputemail">
                          			<option value="false">False</option>
                          			<option value="true">True</option>
                		          </select> <br/>
                  inputCardNumber: <select name="cpinputcardnumber">
                                			<option value="false">False</option>
                                			<option value="true">True</option>
                      		          </select> <br/>
                  inputCVV: <select name="cpinputcvv">
                                <option value="false">False</option>
                                <option value="true">True</option>
                              </select> <br/> </li>
              <li>FormCompletionTime: <input type="text" name="formcompletiontime" value={{.FormCompletionTime}}> </li>


            </ul>
            <input type="submit" value="Submit Changes">
        </form>
    </body>
</html>
