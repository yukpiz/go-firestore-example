var admin = require("firebase-admin");
var acc = require("./yukpiz-labo.json");

admin.initializeApp({
  credential: admin.credential.cert(acc),
  databaseURL: "https://yukpiz-labo.firebaseio.com",
});

var db = admin.firestore();

var query = db.collection("messages").orderBy("SentAt");

var observer = query.onSnapshot(querySnapshot => {
  querySnapshot.docChanges().forEach(change => {
    if (change.type !== "added") {
      return;
    }
    var d = new Date(change.doc.data().SentAt * 1000);
    var h = d.getHours();
    var m = ("0" + d.getMinutes()).substr(-2);
    console.log(`${change.doc.data().SenderID}: ${change.doc.data().Message}  ${h}:${m}`);
  });
}, err => {
  console.log(`Encountered error: ${err}`);
});