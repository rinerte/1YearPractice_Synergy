export const ApiService = {
    async getPages() {
      // const response = await fetch('api.com/pages');
      // if (!response.ok) throw new Error('Failed to fetch pages');
      // return await response.json();
      const response = [{id:1,name:"First"}, {id:2,name:"Second"}, {id:3,name:"Third"}, {id:4,name:"Third"}];
      return response;
    },
  
    async getPageItems(pageId) {
      // const response = await fetch(`api.com/pages/${pageId}`);
      // if (!response.ok) throw new Error('Failed to fetch page items');
      // return await response.json();
      switch(pageId){
        case 1: return [{name:"1"},{name:"1111"},{name:"11"}];
        case 2: return [{name:"2222"},{name:"22"},{name:"2222"}];
        default: return [{name:"3333"},{name:"3333"},{name:"3333"}];
      }
    }
  };