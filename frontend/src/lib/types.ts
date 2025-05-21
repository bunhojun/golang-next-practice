export interface Todo {
    id: string;
    text: string;
    done: boolean;
    user: {
      id: string;
      name: string;
    };
  }
  
  export interface TodosResponse {
    todos: Todo[];
  }

  export interface TodoResponse {
    todo: Todo;
  }