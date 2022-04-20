package entities;

public class Colaboradores {

    Integer id;
    String nameColab;
    Double salaryColab;

    public Colaboradores(Integer id, String nameColab, Double salaryColab){
        this.id = id;
        this.nameColab = nameColab;
        this.salaryColab = salaryColab;
    }

    public Colaboradores() {
    }

    public Integer getId() {
        return id;
    }

    public String getNameColab() {
        return nameColab;
    }

    public Double getSalaryColab() {
        return salaryColab;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public void setNameColab(String nameColab) {
        this.nameColab = nameColab;
    }

    public void setSalaryColab(Double salaryColab) {
        this.salaryColab = salaryColab;
    }

    public void increaseSalary(Double incrementSalary){
        salaryColab += salaryColab * incrementSalary / 100.0;
    }

    @Override
    public String toString() {
        return "List of employees: \n" +
                 id + ", " + nameColab + ", " +salaryColab;
    }
}
