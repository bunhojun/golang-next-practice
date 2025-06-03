import { graphqlClient } from "@/lib/graphql";
import { UserResponse } from "@/lib/types";
import { revalidatePath } from "next/cache";

const GET_USER = `
  query GetUser($id: ID!) {
    user(id: $id) {
      id
      name
      todos {
        id
        text
        done
      }
    }
  }
`;

const CREATE_TODO = `
  mutation CreateTodo($text: String!, $userId: ID!) {
    createTodo(input: { text: $text, userId: $userId }) {
      id
      text
      done
      user {
        id
      }
    }
  }
`

export default async function UserPage(props: { params:  Promise<{ id: string }> }) {
    const params = await props.params;
    const id = params.id;

    const { user } = await graphqlClient.request<UserResponse>(GET_USER, { id });

    return (
        <div>
            <h1>{user.name}</h1>
            <form action={async (formData: FormData) => {
                "use server";
                const text = formData.get("text");
                await graphqlClient.request(CREATE_TODO,
                  {
                    text,
                    userId: id
                  }
                );
                revalidatePath(`/users/${id}`);
            }}>
                <label>
                  <span className="mr-2">Todo</span>
                  <input type="text" name="text" className="border-1 border-gray-300 rounded-md p-2 mr-2" />
                </label>
                <button className="cursor-pointer" type="submit">Create Todo</button>
            </form>
            <ul>
                {user.todos.map((todo) => (
                    <li key={todo.id}>{todo.text}</li>
                ))}
            </ul>
        </div>
    )
}