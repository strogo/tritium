html() {
  # Remove iframe
  ##################
  $(".//iframe") {
    remove()
  }

  # JS Removal
  ##################
  $(".//script[contains(@src,'animation-min.js')]") {
    remove()
  }
  $(".//script[contains(@src,'cacheUtility-min.js')]") {
    remove()
  }
  $(".//script[contains(@src,'cmFunctions.js')]") {
    remove()
  }
  $(".//script[contains(@src,'cmIO-min.js')]") {
    remove()
  }
  $(".//script[contains(@src,'connection-min.js')]") {
    remove()
  }
  $(".//script[contains(@src,'container-min.js')]") {
    remove()
  }
  # Can't remove because it screws up sign-in,out
  # $(".//script[contains(@src,'cookie-beta-min.js')]") {
  #   remove()
  # }
  
  # Can't remove because it screws up sign-in,out
  # $(".//script[contains(@src,'cookie-min.js')]") {
  #   remove()
  # }
  
  # Can't remove -- this breaks signin redirection
  # $(".//script[contains(@src,'cookieJar.js')]") {
  #   remove()
  # }
  
  # Coremetrics Stuff
  
  # Yupu regexes this one
  # $(".//script[contains(@src,'coremetrics.cmdatatagutils.js')]") {
  #   remove()
  # }
  # $(".//script[contains(@src,'coremetrics.eluminate.js')]") {
  #   remove()
  # }
  # $(".//script[contains(@src,'coremetrics.io_v4.js')]") {
  #   remove()
  # }
  $(".//script[contains(@src,'coremetrics.techprops.js')]") {
    remove()
  }
  $(".//script[contains(@src,'exception-min.js')]") {
    remove()
  }
  $(".//script[contains(@src,'flashUtils.js')]") {
    remove()
  }
  
  # Used to validate the search input
  $(".//script[contains(@src,'formValidation.js')]") {
    remove()
  }
  
  # Can't remove because it screws up sign-in,out
  # $(".//script[contains(@src,'globalCart-min.js')]") {
  #   remove()
  # }
  $(".//script[contains(@src,'globalHeaderPool.js')]") {
    remove()
  }
  
  # Can't Remove because it screws up the cart item number
  # $(".//script[contains(@src,'macys-min.js')]") {
  #   remove()
  # }
  $(".//script[contains(@src,'mbox.js')]") {
    remove()
  }
  $(".//script[contains(@src,'quickbag-min.js')]") {
    remove()
  }
  $(".//script[contains(@src,'rules_90067660.js')]") {
    remove()
  }
  
  # I haven't found this
  # $(".//script[contains(@src,'signintest2.js')]") {
  #   remove()
  # }
  $(".//script[contains(@src,'standard.js')]") {
    remove()
  }
  $(".//script[contains(@src,'urlUtility-min.js')]") {
    remove()
  }
  $(".//script[contains(@src,'userObject.js')]") {
    remove()
  }
  # Can't remove because it screws up sign-in,out
  # $(".//script[contains(@src,'yahoo-dom-event.js')]") {
  #   remove()
  # }
  $(".//script[contains(@src,'mbox/standard')]") {
    remove()
  }
  $(".//script[contains(.,'1@x96')]") {
    remove()
  }
}
