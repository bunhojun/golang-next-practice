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

  export interface User {
    id: string;
    name: string;
    todos: Todo[];
  }
  
  export interface UsersResponse {
    users: User[];
  }

  export interface UserResponse {
    user: User;
  }

  export interface CreateTodoResponse {
    createTodo: Todo;
  }

  export interface CreateTodoInput {
    text: string;
    userId: string;
  }
