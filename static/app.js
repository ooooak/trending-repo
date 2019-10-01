// sort items by star
const flatten = function (storage){
  let ret = [];
  for (const key of Object.keys(storage)) {
    let local = storage[key];
    if (local !== null && local !== void 0){      
      ret = ret.concat(local)
    }
  }
  return ret;
}

const sortCallback = function(a, b){
  return a.stars > b.stars ? -1 : 1;
}

const dateList = function(storage){
  let ret = {};
  for (const key of Object.keys(storage).splice(-20)) {
    let local = storage[key];    
    if (local == null || local == void 0){
      local = [];
    }
    ret[key] = local.sort(sortCallback).splice(0, 50);
    // .filter(item => item.language == "Clojure")
  }
  return ret;
};

(function () {
  const storage = __output;
  __output = {};

  // __output["All"] = flatten(storage).sort(sortCallback);
  // console.log(__output);  
  __output = dateList(storage);
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