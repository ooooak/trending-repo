// sort items by star
(function () {
  const storage = __output;
  __output = {};

  for (const key of Object.keys(storage).splice(-20)) {
    let data = storage[key];    
    if (data == null || data == void 0){
      data = [];
    }

    

    __output[key] = data.sort((a, b) => {
      return a.stars > b.stars ? -1 : 1;
    })
    // .filter(item => item.language == "Clojure")
    .splice(0, 50);;
  }
})();

const buildDate = function(date){
  const [d, m, y] = date.split("-")
  return new Date(y, m - 1, d);
};

var app = new Vue({
  el: '#app',
  data: {
    // items: __output,
  },
  computed:{
    itemsKeys: function(){
      return Object.keys(__output).sort((a, b) => {
        return buildDate(a) > buildDate(b) ? -1 : 1;
      });
    }, 
    items: function(){
      return __output
    }
  }
})