// sort items by star
(function () {  
  for (const key in __output) {
    __output[key] = __output[key].sort((a, b) => {
      return a.stars > b.stars ? -1 : 1;
    });
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