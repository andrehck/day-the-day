package com.company;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        int n = sc.nextInt();
        int m = sc.nextInt();

        int[][] matrix = new int[n][m];

        for (int i =0; i <matrix.length; i++){
            for (int j=0;j<n ; j++){
                matrix[i][j] = sc.nextInt();
            }
        }

        int x = sc.nextInt();

        for (int i =0; i <matrix.length; i++){
            for (int j=0;j<n ; j++){
                if (matrix[i][j] == x){
                    System.out.println("Position: " + i + ", " + j );
                }
            }
        }
    }
}
