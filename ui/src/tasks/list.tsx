import React, { useState } from 'react';
import { useQuery, useMutation } from '../hooks';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Paper,
  TextField,
  Button
} from '@material-ui/core';

interface GetTasks {
  tasks: {
    data: Array<{
      id: number;
      title: string;
      description: string;
      priority: number;
    }>
  }
}

const GET_TASKS = `
  query {
    tasks {
      data {
        id
        title
        description
        priority
      }
    }
  }
`;

interface CreateTask {
  data: {
    createTask: {
      id: number;
    }
  }
}

const CREATE_TASK = `
  mutation($title: String!, $description: String!, $priority: Int!) {
    createTask(title: $title, description: $description, priority: $priority) {
      id
    }
  }
`;

interface UpdateTask {
  data: {
    updateTask: {
      id: number;
    }
  }
}

const UPDATE_TASK = `
  mutation($id: Int!, $title: String!, $description: String!, $priority: Int!) {
    updateTask(id: $id, title: $title, description: $description, priority: $priority) {
      id
    }
  }
`;

interface DeleteTask {
  data: {
    deleteTask: {
      id: number;
    }
  }
}

const DELETE_TASK = `
  mutation($id: Int!) {
    deleteTask(id: $id) {
      id
    }
  }
`;

const API_URI = 'http://localhost:8081';

function TaskList() {
  const { data, refetch } = useQuery<GetTasks, {}>(API_URI, GET_TASKS, {});
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [priority, setPriority] = useState<number>(0);

  const { run: createTask } = useMutation<CreateTask, {}>(API_URI, CREATE_TASK, {
    variables: { title, description, priority }
  });

  const { run: updateTask } = useMutation<UpdateTask, {}>(API_URI, UPDATE_TASK, {});


  const { run: deleteTask } = useMutation<DeleteTask, {}>(API_URI, DELETE_TASK, {});

  const onAdd = async () => {
    await createTask();
    await refetch();
    setTitle('');
    setDescription('');
    setPriority(0);
  };

  const onRemove = async (id: number) => {
    await deleteTask({ variables: { id } });
    await refetch();
  };

  const onUpdate = async (task: GetTasks['tasks']['data'][0]) => {
    await updateTask({ variables: { ...task } });
    await refetch();
  };

  const rows = data && data.tasks.data.map((task) => (
    <TableRow key={ task.id }>
      <TableCell>
        <TextField
          placeholder="Title"
          defaultValue={ task.title }
          onChange={ ({ currentTarget: { value } }) => onUpdate({ ...task, title: value }) }
        />
      </TableCell>
      <TableCell>
        <TextField
          placeholder="Description"
          defaultValue={ task.description }
          onBlur={ ({ currentTarget: { value } }) => onUpdate({ ...task, description: value }) }
        />
      </TableCell>
      <TableCell>
        <TextField
          type="number"
          placeholder="Priority"
          defaultValue={ task.priority || 0 }
          onChange={ ({ currentTarget: { value } }) => onUpdate({ ...task, priority: value ? parseInt(value) : 0 }) }
        />
      </TableCell>
      <TableCell><Button onClick={ () => onRemove(task.id) }>Remove</Button></TableCell>
    </TableRow>
  ));

  return (
    <>
      <Paper>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>
                <TextField
                  placeholder="Title"
                  value={ title }
                  onChange={ ({ currentTarget: { value } }) => setTitle(value) }
                />
              </TableCell>
              <TableCell>
                <TextField
                  placeholder="Description"
                  value={ description }
                  onChange={ ({ currentTarget: { value } }) => setDescription(value) }
                />
              </TableCell>
              <TableCell>
                <TextField
                  type="number"
                  placeholder="Priority"
                  value={ priority || 0 }
                  onChange={ ({ currentTarget: { value } }) => setPriority(value ? parseInt(value, 10) : 0) }
                />
              </TableCell>
              <TableCell>
                <Button onClick={ onAdd }>Add</Button>
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            { rows }
          </TableBody>
        </Table>
      </Paper>
    </>
  );
}

export { TaskList };
