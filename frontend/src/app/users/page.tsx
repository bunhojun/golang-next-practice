
import { graphqlClient } from "@/lib/graphql";
import { UsersResponse } from "@/lib/types";
import Link from "next/link";

const GET_USERS = `
  query {
    users {
      id
      name
    }
  }
`


export default async function UsersPage() {
    const { users } = await graphqlClient.request<UsersResponse>(GET_USERS);

    return (
        <div>
            <h1>Users</h1>
            <ul>
                {users.map((user) => (
                    <li key={user.id}>
                        <Link href={`/users/${user.id}`}>
                            {user.name}
                        </Link>
                    </li>
                ))}
            </ul>
        </div>
    )
}
