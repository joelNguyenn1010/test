const data = require('./students.json');

const totalStudent = (students) => {

    let returnValue = [];

    // An object to capture the lenght of the class to calculate the average student
    let repeatter = {};

    students.map((student) => {

        let shouldPush = true;
        returnValue.map((value, index) => {
            if(value[student['class']]) {
                shouldPush = false;

                returnValue[index][student['class']]['total'] += student.attended;

                repeatter[student['class']] += 1;

                returnValue[index][student['class']]['average'] =  
                Math.round(
                    returnValue[index][student['class']]['total'] / repeatter[student['class']]
                )
            } 
        })

        if(shouldPush) {

            repeatter[student['class']] = 1;

            returnValue.push({
                [student['class']]: {
                    total: student.attended,
                    average: Math.round(student.attended/repeatter[student['class']])
                }
            })
        }
    });

    return returnValue;
}


// Assumption
// attended: Integer
// students: Array
console.log(totalStudent(data.students));