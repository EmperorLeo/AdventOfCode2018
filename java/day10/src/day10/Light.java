package day10;

import javafx.scene.paint.Color;
import javafx.scene.shape.Rectangle;

public class Light extends Rectangle {
	private int x;
	private int y;
	private int velocityX;
  private int velocityY;


	public Light(int x, int y, int velocityX, int velocityY) {
    super(1, 1, Color.DARKRED);
		this.x = x;
		this.y = y;
		this.velocityX = velocityX;
		this.velocityY = velocityY;
	}
	
	public int getPropertyX() {
		return this.x;
	}
	
	public int getPropertyY() {
		return this.y;
	}
		
	public void travel() {
		x += velocityX;
		y += velocityY;
		// this.setX(x);
		// this.setY(y);
		
		moveLightOnScreen();
	}
	
	public void setPerspective() {
		moveLightOnScreen();
	}
	
	private void moveLightOnScreen() {
		this.setLayoutX(this.x);
		this.setLayoutY(this.y);
	}
		
}
