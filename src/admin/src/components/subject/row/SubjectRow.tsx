import { useState, useEffect } from 'react';
import axios from 'axios';
import { Checkbox } from '@mui/material';
import { API_URL } from 'App';

interface Subject {
  IsAllowed: boolean;
  SubjectID: number;
}

interface User {
  ID: number;
  Name: string;
}

type SubjectRowProps = {
  user: User;
  subject: Subject;
  roleID: number;
};

export default function SubjectRow({ user, subject, roleID }: SubjectRowProps) {
  //   const label = { inputProps: { 'aria-label': 'Checkbox demo' } };
  const [checked, setChecked] = useState(subject.IsAllowed);

  useEffect(() => {
    setChecked(subject.IsAllowed);
  }, [roleID, subject]);

  const addSubjectAssignment = async (subjectID: number, roleID: number) => {
    const data = {
      SubjectID: subjectID,
      RoleID: roleID,
    };
    await axios.post(`${API_URL}/rbac/subject-assignment`, data);
  };
  const deleteSubjectAssignment = async (subjectID: number, roleID: number) => {
    await axios.delete(
      `${API_URL}/rbac/subject-assignment?subjectID=${subjectID}&roleID=${roleID}`,
    );
  };

  const handleChange = (
    event: React.ChangeEvent<HTMLInputElement>,
    roleID: number,
    subjectID: number,
  ) => {
    setChecked(event.target.checked);
    if (event.target.checked === true) {
      addSubjectAssignment(subjectID, roleID);
    } else {
      deleteSubjectAssignment(subjectID, roleID);
    }
  };

  return (
    <tr key={subject.SubjectID} style={{ height: '55px' }}>
      <td style={{ textAlign: 'center' }}>{subject.SubjectID}</td>
      <td style={{ textAlign: 'center' }}>{user.Name}</td>
      <td style={{ textAlign: 'center' }}>
        <Checkbox
          checked={checked}
          onChange={(e) => handleChange(e, roleID, subject.SubjectID)}
        />
      </td>
    </tr>
  );
}
