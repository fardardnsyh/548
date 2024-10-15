function getChkBoxValue() {
    /*alert("Hello World from Get Check Box Value !")*/
    chkbox1 = document.getElementById("chkbox1")
    chkbox2 = document.getElementById("chkbox2")
    chkbox3 = document.getElementById("chkbox3")
    chkbox4 = document.getElementById("chkbox4")
    chkbox5 = document.getElementById("chkbox5")
    numberChecked = 0
    chkbox1Val = chkbox1.value
    chkbox2Val = chkbox2.value
    chkbox3Val = chkbox3.value
    chkbox4Val = chkbox4.value
    chkbox5Val = chkbox5.value
    if (chkbox1.checked) {
       numberChecked = numberChecked + 1
    }
    if (chkbox2.checked) {
      numberChecked = numberChecked + 1
    }
    if (chkbox3.checked) {
      numberChecked = numberChecked + 1
    }
    if (chkbox4.checked) {
      numberChecked = numberChecked + 1
    }
    if (chkbox5.checked) {
      numberChecked = numberChecked + 1
    }
    /*alert (`Check Box 1: ${chkbox1.checked} ; Value: ${chkbox1Val}`);
    alert (`Check Box 2: ${chkbox2.checked} ; Value: ${chkbox2Val}`);
    alert (`Check Box 3: ${chkbox3.checked} ; Value: ${chkbox3Val}`);
    alert (`Check Box 4: ${chkbox4.checked} ; Value: ${chkbox4Val}`);
    alert (`Check Box 5: ${chkbox5.checked} ; Value: ${chkbox5Val}`);*/
    alert (`Number of Checked Boxes: ${numberChecked}`);
    if (numberChecked == 0) {
       alert("Error: You have NOT selected any Category;\nYou must select at least 1 Category.")    
    }/*if*/
  }
  function getChkBox1Value() {
    chkbox1 = document.getElementById("chkbox1")
    /*alert (`Check Box 1: ${chkbox1.checked}`)*/
  }
  function getChkBox2Value() {
    chkbox2 = document.getElementById("chkbox2")
    /*alert (`Check Box 2: ${chkbox2.checked}`)*/
  }
  function getChkBox3Value() {
    chkbox3 = document.getElementById("chkbox3")
    /*alert (`Check Box 3: ${chkbox3.checked}`)*/
  }
  function getChkBox4Value() {
    chkbox4 = document.getElementById("chkbox4")
    /*alert (`Check Box 4: ${chkbox4.checked}`)*/
  }
  function getChkBox5Value() {
    chkbox5 = document.getElementById("chkbox5")
    /*alert (`Check Box 5: ${chkbox5.checked}`)*/
  }

// Shows/Hides Comment Section
